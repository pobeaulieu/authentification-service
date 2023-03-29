package controllers

import (
	"backend/database"
	"backend/models"
	"backend/utility"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func User(c *fiber.Ctx) error {
	token, err := VerifyAuthentification(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("user", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := getUserFromToken(token)

	utility.LogInfo(user.Name, "user")

	return c.JSON(user)
}

func Users(c *fiber.Ctx) error {
	token, err := VerifyAuthentification(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("user", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := getUserFromToken(token)

	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		utility.LogInfo("user", "noUsers")
		return c.JSON(fiber.Map{
			"message": "Aucun utilisateurs",
		})
	}

	// if doesnt have admin or business role, return an error
	if user.AdminRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid admin role",
		})
	}

	utility.LogInfo(user.Name, "users")

	return c.JSON(users)
}

func getUserFromToken(token *jwt.Token) *models.User {
	// Conversion de type claims à StandardClaims pour avoir le Issuer
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	// get user
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return &user
}

func getPasswordPolicy(user models.User) *models.PasswordPolicy {
	// Conversion de type claims à StandardClaims pour avoir le Issuer
	var policy models.PasswordPolicy
	// get user
	database.DB.Where("id = ?", user.PasswordPolicyId).First(&policy)

	return &policy
}

func getLoginPolicy(user models.User) *models.LoginPolicy {
	// Conversion de type claims à StandardClaims pour avoir le Issuer
	var policy models.LoginPolicy
	// get user
	database.DB.Where("id = ?", user.LoginPolicyId).First(&policy)

	return &policy
}
