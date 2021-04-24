package main

import (
	"gorm-test/controller"
	"gorm-test/model"
	"gorm-test/repository"
	"gorm-test/services"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	server := echo.New()
	dsn := "host=ec2-3-211-37-117.compute-1.amazonaws.com user=kajzsajbioqcjz password=7273573c396f3e10b1510677a30bd2758b906c16c5abb11e0e4742bc10cb6e5c dbname=dge6vvn2lvfgs port=5432"
	db, _ := gorm.Open(postgres.Open(dsn))
	c := controller.New(services.New(repository.New(db)))
	db.AutoMigrate(&model.People{})
	db.AutoMigrate(&model.Pet{})
	server.GET("/peoples", c.GetPeoples)
	server.POST("/peoples", c.PostPeople)
	server.Start(":80")
}
