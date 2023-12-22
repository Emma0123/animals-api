package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"

	"animals-api/animal"
	"animals-api/handlers"

	"gorm.io/gorm"
)

func main() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB CONNECTION ERROR")
	}

	// db.AutoMigrate(&animal.RequestAnimals{})
	db.AutoMigrate(&animal.Animals{})

	animalRepository := animal.NewRepository(db)

	animalService := animal.NewService(animalRepository)

	animalHandler := handlers.NewAnimalHandler(animalService)

	//============

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/animal", animalHandler.GetAllAnimalsDetail)
	v1.GET("/animal/:id", animalHandler.GetOneAnimalDetail)
	v1.POST("/animal/create-animal", animalHandler.CreateAnimal)
	v1.PUT("/animal/:id", animalHandler.UpdateAnimal)
	v1.DELETE("/animal/:id", animalHandler.DeleteAnimalDetail)

	router.Run()
}

// main
// handler
// service => terkait dengan logic eg. bisa upload product
// repository
// db
// mysql
