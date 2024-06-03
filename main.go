package main

import (
	"github.com/gin-gonic/gin"

	"restfull-api/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main() {
	dsn := "host=localhost user=dev password=qwerty dbname=go-restfull-api port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	  }

	router := gin.Default()
	routes.SetupRouter(router)
	
	router.Run() 
}