package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var router *gin.Engine

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

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
	fmt.Println("Webhook received...")
	log.Infoln("Test -  Info")
	log.Warnln("Test - Warning")
	log.WithFields(log.Fields{
		"GCP Project":  "Test -  Project",
		"Product Name": "Test - Product",
		"size":         10}).Info("Test - Infoln")
	log.Errorln("Test - Error")
	c.String(http.StatusOK, "POST request received...")
}
func fetchHelloWorld(c *gin.Context) {
	fmt.Println("received GET call...")
	log.Infoln("Test -  Info")
	log.Warnln("Test - Warning")
	log.WithFields(log.Fields{
		"GCP Project":  "Test -  Project",
		"Product Name": "Test - Product",
		"size":         10}).Info("Test - Infoln")
	log.Errorln("Test - Error")
	c.JSON(http.StatusOK, "Hello World !!")
}
