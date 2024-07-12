package main

import (
	"log"
	"os"

	"github.com/adedeji/resturant-management-system/database"
	"github.com/adedeji/resturant-management-system/middleware"
	"github.com/adedeji/resturant-management-system/routes"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

var resturantCollection *mongo.Collection = database.OpenCollection(database.Client, "resturant")

func main() {
	if err := gotdotenv.Load; err != nil {
		log.Fatal("NO env Found")

	}

	port := os.Getenv("PORT")

	if port == "" {
		port = 4000
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())

	routes.FoodRoutes()
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.NoteRoute(router)
	routes.OrderRoute(router)
	routes.TableRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)

}
