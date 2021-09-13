package controllers

import (
	"time"

	"github.com/KinitaL/golang-api-v2/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	db := connectToDB()
	db.AutoMigrate(&models.User{})
	var user models.User
	c.BodyParser(&user)
	db.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	db := connectToDB()
	db.AutoMigrate(&models.User{})

	//data from user
	var user models.User
	c.BodyParser(&user)

	//data from DB
	var dbuser models.User
	db.Where("login = ?", user.Login).First(&dbuser)

	//compare
	if user.Password == dbuser.Password {
		cookie := new(fiber.Cookie)
		cookie.Name = "token"
		cookie.Value = "auth-token"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.Cookie(cookie)
		return c.JSON("You're logged in")
	} else {
		return c.JSON("wrong auth data")
	}
}

func DeleteUser(c *fiber.Ctx) error {
	db := connectToDB()
	db.AutoMigrate(&models.User{})
	var user models.User
	c.BodyParser(&user)
	db.Where("login = ?", &user.Login).Delete(&user)
	return c.JSON(user.Login + " was deleted")
}
