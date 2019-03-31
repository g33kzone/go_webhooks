package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	initializeRoutes()
	router.Run("0.0.0.0:8080")
}

func initializeRoutes() {
	router.POST("/api", handleWebHooks)
	router.GET("/api", fetchHelloWorld)
}

func handleWebHooks(c *gin.Context) {
	fmt.Print("Webhook received...")
	c.String(http.StatusOK, "POST request received...")
}
func fetchHelloWorld(c *gin.Context) {
	fmt.Print("received GET call...")
	c.JSON(http.StatusOK, "Hello World !!")
}
