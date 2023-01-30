package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("starting on :%s\n", port)
	r.Run(":" + port)
}
