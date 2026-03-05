package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

)

func StartAPI(Port int) {
	fmt.Println("[+] Starting API")

	router := gin.Default()

	router.LoadHTMLGlob("api/v1/templates/root.tmpl")
	router.StaticFile("/assets/fenris.png", "./assets/fenris.png")
	router.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK, "root.tmpl", gin.H{
			"title": "Fenris",
		})
	})
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", V1)

		env := v1.Group("/env")
		env.GET("/", GetEnv)

		db := v1.Group("/database")
		db.GET("/", Placeholder("/database", "Starting..."))

		//TODO: Implement config page
		config := v1.Group("/config")
		config.GET("/", Placeholder("/config", "Starting..."))

		file := v1.Group("/file")
		file.GET("/", GetFiles)

		time := v1.Group("/time")
		time.GET("/", GetTime)

		test := v1.Group("/test")
		test.GET("/", Placeholder("/", "fenris didn't have time to implement this"))
	}
	router.Run(fmt.Sprintf(":%d", Port))
}
