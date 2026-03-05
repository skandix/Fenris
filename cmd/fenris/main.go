package main

import (
	"fmt"
	"os"

	"fenris/api/v1"
	"fenris/config"
	"fenris/database"
	"fenris/cache"
)

// Main
func main() {
	REDIS_HOST := os.Getenv("REDIS_HOST")

	fmt.Println("=== Fenris ===")
	cache.ConnectToRedis(REDIS_HOST, "6379")
	database.StartDatabase()
	config.StartConfig()
	api.StartAPI(8080)
}