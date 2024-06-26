package main

import (
	"github.com/atanda0x/e-commerce/database"
	"github.com/atanda0x/e-commerce/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	database.DBset()
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use()

	router.GET("/addToCart")
	router.GET("/removeItem")
	router.GET("/cartCheckout")
	router.GET("/instantBuy")
}
