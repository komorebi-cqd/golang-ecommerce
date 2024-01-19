package main

// "github.com/komoreb-cqd/golang-ecommerce/controllers"
// "github.com/komoreb-cqd/golang-ecommerce/database"
// "github.com/komoreb-cqd/golang-ecommerce/middlewarw"
// "github.com/komoreb-cqd/golang-ecommerce/routes"
// "github.com/gin-gonic/gin"

import (
	"komorebi-ecommerce/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	// router.Use()

	// router.GET("/addtocart")
}
