package controllers

import (
	"log"
	"os"

	"github.com/KinitaL/golang-api-v2/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Get(c *fiber.Ctx) error {
	db := connectToDB()
	db.AutoMigrate(&models.Product{})
	var products []models.Product
	db.Find(&products)
	return c.JSON(products)
}

func Post(c *fiber.Ctx) error {
	db := connectToDB()
	db.AutoMigrate(&models.Product{})
	var product models.Product
	c.BodyParser(&product)
	db.Create(&product)
	return c.JSON(product)
}

func Put(c *fiber.Ctx) error {
	db := connectToDB()
	db.AutoMigrate(&models.Product{})
	var product models.Product
	c.BodyParser(&product)
	db.Model(models.Product{}).Where("id=?", c.Params("id")).Updates(&product)
	return c.JSON(product)
}

func Delete(c *fiber.Ctx) error {
	db := connectToDB()
	db.AutoMigrate(&models.Product{})
	var product models.Product
	c.BodyParser(&product)
	db.Where("name = ?", &product.Name).Delete(&product)
	return c.JSON(product)
}

func envLoad() {
	if err := godotenv.Load("./db.env"); err != nil {
		log.Fatal("No .env file found")
	}
}

func envManage(key string) string {

	value, exists := os.LookupEnv(key)
	if exists {
		return value
	} else {
		log.Fatal("No such env")
		return ""
	}
}

func connectToDB() *gorm.DB {
	envLoad()
	host := envManage("POSTGRES_HOST")
	port := envManage("POSTGRES_PORT")
	user := envManage("POSTGRES_USER")
	pass := envManage("POSTGRES_PASSWORD")
	dbname := envManage("POSTGRES_DB")
	helper := "host=" + host + " port=" + port + " user=" + user + " password=" + pass + " dbname=" + dbname + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(helper), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
