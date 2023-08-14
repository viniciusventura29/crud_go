package main

import (
	"crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.AppRoutes(app)

	app.Run("localhost:4000")
}
