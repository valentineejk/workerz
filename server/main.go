package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/valentineejk/workerz/controller"
	"github.com/valentineejk/workerz/database"
	"github.com/valentineejk/workerz/repository"
	"github.com/valentineejk/workerz/router"
	"github.com/valentineejk/workerz/services"
)

func main() {

	//load env
	env := LoadEnv()
	if env != nil {
		fmt.Println("Error loading environment:", env)
		return
	}

	//gin
	g := gin.Default()
	gin.Logger()

	//db
	conn := database.ConnectToDB()
	// Initialize gepository
	repo := repository.NewDbWorkerzRepo(conn)
	// Initialize service
	workerzService := services.NewWorkerzService(repo)
	// Initialize controller
	workerzController := controller.NewWorkerzController(workerzService)

	//route
	router.Router(g, workerzController)

	PORT := os.Getenv("PORT")
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Running
	g.Run(":" + PORT)

}
