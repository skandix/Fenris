package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"fenris/files"
)

// root routee
func Placeholder(RoutePath string, RouteName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			fmt.Sprintf("%s", RoutePath): fmt.Sprintf("%s", RouteName),
		})
	}
}

func V1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"/v1/": "Online",
		"/v1/env": "Online",
		"/v1/file": "Online",
		"/v1/database": "Offline",
		"/v1/time": "Online",
		"/v1/config": "Offline",
	})
}

// Read out environment variable, and show on a route
func GetEnv(c *gin.Context) {
	tenv := "BEST_MOVIE"
	if (os.Getenv(tenv) == "HACKERS") {
		c.JSON(http.StatusOK, gin.H{
			tenv: os.Getenv(tenv),
		})
	} else {
		panic(fmt.Sprintf("[-] expected value in env %s not found", tenv))
	}
}

func GetFiles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"file": files.GetFile(".secret"),
	})
}

// TODO: Needs to be implemented
func GetDatabase() {
}

// TODO: Needs to be implemented
func WriteDatabase() {

}

// Get the time
func GetTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"time": time.Now(),
	})
}

// TODO: Needs to be implemented
func GetConfig() {

}

// Dummy route
func TestRoutes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
