package controllers

import (
	"backend/database"
	"backend/models"
	"backend/utility"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var pwdpolicy models.PasswordPolicy
	database.DB.Where("id = ?", 1).First(&pwdpolicy)

	err := pwdpolicy.Validate(data["password"], nil)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	adminRole, _ := strconv.Atoi(data["adminRole"])
	businessRole, _ := strconv.Atoi(data["businessRole"])
	residentialRole, _ := strconv.Atoi(data["residentialRole"])

	user := models.User{
		Name:              data["name"],
		Email:             data["email"],
		Password:          password,
		AdminRole:         uint(adminRole),
		BusinessRole:      uint(businessRole),
		ResidentialRole:   uint(residentialRole),
		Blocked:           0,
		LastModified:      time.Now(),
		LastLoginAttempt:  time.Time{},
		LoginAttemptCount: 0,
	}

	err = database.DB.Create(&user).Error

	if err != nil {
		utility.LogInfo("system", "errCreateUser")
		return c.JSON(fiber.Map{
			"message": "error creating the user. Try with another email. ",
			"error":   err.Error(),
		})
	}

	utility.LogInfo(data["name"], "createUser")

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		utility.LogInfo("system", "userNotFound")
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	loginPolicy := getLoginPolicy(user)

	if time.Since(user.LastLoginAttempt) < time.Duration(loginPolicy.LoginTimeInterval)*time.Second {
		// L'utilisateur a dépassé la limite de tentatives de connexion et il n'a pas encore attendu suffisamment longtemps
		utility.LogInfo("system", "shortTimeInterval")
		return c.JSON(fiber.Map{
			"message": "too short time interval between attempts",
		})
	}

	if user.Blocked == 1 || user.LoginAttemptCount > loginPolicy.MaxLoginAttemptCount {
		c.Status(fiber.StatusBadRequest)
		// Incrementer le nombre de tentatves manquees
		user.LoginAttemptCount = user.LoginAttemptCount + 1
		user.Blocked = 1
		// Mettre a jour le last login manque
		user.LastLoginAttempt = time.Now()
		database.DB.Save(&user)
		utility.LogInfo("system", "accountBlocked")
		return c.JSON(fiber.Map{
			"message": "The account is blocked. Please contact an administrator to unblock the account.",
		})

	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		// Incrementer le nombre de tentatves manquees
		user.LoginAttemptCount = user.LoginAttemptCount + 1
		user.LastLoginAttempt = time.Now()
		database.DB.Save(&user)
		utility.LogInfo("system", "wrongPassword")
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), //JWT valide 1 heure
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRETKEY")))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		utility.LogInfo("system", "failedLogin")
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	// Reset le compteur d'attempts
	user.LoginAttemptCount = 0
	user.LastLoginAttempt = time.Time{}

	// Save the changes to the user object
	if err := database.DB.Save(&user).Error; err != nil {
		fmt.Println("Error in database save:", err)
		c.Status(fiber.StatusInternalServerError)
		utility.LogInfo("system", "errorDBSave")
		return c.JSON(fiber.Map{
			"message": "could not save user data",
		})
	}

	utility.LogInfo(user.Name, "login")

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Logout(c *fiber.Ctx) error {
	// creer un cookie dans le passé pour "retirer" le cookie du browser
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	utility.LogInfo("system", "logout")

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func VerifyAuthentification(c *fiber.Ctx) (*jwt.Token, error) {
	// if not identified, return an error
	cookie := c.Cookies("jwt")

	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRETKEY")), nil
	})
}
