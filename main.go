package main

import (
	"github.com/ashkan-developer/ecommerce-golang/controllers"
	"github.com/ashkan-developer/ecommerce-golang/databasea"
	"github.com/ashkan-developer/ecommerce-golang/middleware"
	"github.com/ashkan-developer/ecommerce-golang/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(databasea.ProductData(databasea.Client, "Products"), databasea.UserData(databasea.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("cartcheckout", app.BuyFromCart())
	router.GET("instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
