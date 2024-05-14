package main

import (
	"github.com/gin-gonic/gin"

	"restfull-api/routes"
)


func main() {
	

	router := gin.Default()
	routes.SetupRouter(router)
	
	router.Run() 
}