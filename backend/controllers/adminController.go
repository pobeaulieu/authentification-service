package controllers

import (
	"backend/database"
	"backend/models"
	"backend/utility"

	"github.com/gofiber/fiber/v2"
)

func Unblock(c *fiber.Ctx) error {
	var data map[string]int

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id := data["id"]

	token, err := VerifyAuthentification(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("admin", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := getUserFromToken(token)

	// if doesnt have admin or business role, return an error
	if user.AdminRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid admin role",
		})
	}

	var userToUnblock models.User
	// get user
	database.DB.Where("id = ?", id).First(&userToUnblock)

	userToUnblock.Blocked = 0
	userToUnblock.LoginAttemptCount = 0

	database.DB.Save(&userToUnblock)

	utility.LogInfo(user.Name, "unblocked")

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func MaxLoginAttempts(c *fiber.Ctx) error {
	var data map[string]int

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	newMax := data["max"]

	if newMax <= 1 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("system", "maxAttemptsLimit")
		return c.JSON(fiber.Map{
			"message": "max attempts must be greater than 1",
		})
	}

	token, err := VerifyAuthentification(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("admin", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := getUserFromToken(token)

	// if doesnt have admin or business role, return an error
	if user.AdminRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid admin role",
		})
	}

	if err := database.DB.Model(&models.LoginPolicy{}).Where("id >= ?", 0).Update("MaxLoginAttemptCount", newMax).Error; err != nil {
		c.Status(fiber.StatusInternalServerError)
		utility.LogInfo("system", "errMaxAttemptsSave")
		return c.JSON(fiber.Map{
			"message": "could not save the new max",
		})
	}

	utility.LogInfo(user.Name, "maxAttempts")

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func LoginTimeInterval(c *fiber.Ctx) error {
	var data map[string]int

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	time := data["seconds"]

	if time <= 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("system", "minLoginTimeInterval")
		return c.JSON(fiber.Map{
			"message": "time interval must be positive",
		})
	}

	token, err := VerifyAuthentification(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("admin", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := getUserFromToken(token)

	// if doesnt have admin or business role, return an error
	if user.AdminRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid admin role",
		})
	}

	if err := database.DB.Model(&models.LoginPolicy{}).Where("id >= ?", 0).Update("LoginTimeInterval", time).Error; err != nil {
		c.Status(fiber.StatusInternalServerError)
		utility.LogInfo("system", "errMinLoginTimeIntervalSave")
		return c.JSON(fiber.Map{
			"message": "could not save the new time interval",
		})
	}

	utility.LogInfo(user.Name, "loginTimeInterval")

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func PasswordPolicy(c *fiber.Ctx) error {
	token, err1 := VerifyAuthentification(c)

	if err1 != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("admin", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := getUserFromToken(token)

	// if doesnt have admin or business role, return an error
	if user.AdminRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid admin role",
		})
	}

	var newPolicy models.PasswordPolicy
	err := c.BodyParser(&newPolicy)

	if err != nil {
		return err
	}

	var policy models.PasswordPolicy
	// get user
	database.DB.Where("id = ?", 1).First(&policy)

	if newPolicy.MinLength != nil {
		policy.MinLength = newPolicy.MinLength
	}
	if newPolicy.RequireLower != nil {
		policy.RequireLower = newPolicy.RequireLower
	}
	if newPolicy.RequireUpper != nil {
		policy.RequireUpper = newPolicy.RequireUpper
	}
	if newPolicy.RequireSymbol != nil {
		policy.RequireSymbol = newPolicy.RequireSymbol
	}
	if newPolicy.RequireNumber != nil {
		policy.RequireNumber = newPolicy.RequireNumber
	}
	if newPolicy.HistorySize != nil {
		policy.HistorySize = newPolicy.HistorySize
	}

	database.DB.Save(&policy)

	utility.LogInfo(user.Name, "passwordPolicy")

	return c.JSON(fiber.Map{
		"message": "success",
		"policy":  policy,
	})
}

func GetPasswordPolicy(c *fiber.Ctx) error {

	token, err := VerifyAuthentification(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("admin", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := getUserFromToken(token)

	// if doesnt have admin or business role, return an error
	if user.AdminRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid admin role",
		})
	}

	var policy models.PasswordPolicy
	// get user
	database.DB.Where("id = ?", 1).First(&policy)

	utility.LogInfo(user.Name, "getPasswordPolicy")

	return c.JSON(policy)
}

func GetLoginPolicy(c *fiber.Ctx) error {

	token, err := VerifyAuthentification(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("admin", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := getUserFromToken(token)

	// if doesnt have admin or business role, return an error
	if user.AdminRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid admin role",
		})
	}

	var policy models.LoginPolicy
	// get user
	database.DB.Where("id = ?", 1).First(&policy)

	utility.LogInfo(user.Name, "getLoginPolicy")

	return c.JSON(policy)
}

func UserRole(c *fiber.Ctx) error {
	token, err1 := VerifyAuthentification(c)

	if err1 != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("admin", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := getUserFromToken(token)

	// if doesnt have admin or business role, return an error
	if user.AdminRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid admin role",
		})
	}

	var data map[string]uint

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})

	}

	var userToEdit models.User
	// get user
	database.DB.Where("id = ?", data["userid"]).First(&userToEdit)

	userToEdit.AdminRole = data["adminRole"]
	userToEdit.BusinessRole = data["businessRole"]
	userToEdit.ResidentialRole = data["residentialRole"]

	database.DB.Save(&userToEdit)

	utility.LogInfo(user.Name, "userRole")

	return c.JSON(fiber.Map{
		"message":    "success",
		"editeduser": userToEdit,
	})
}
