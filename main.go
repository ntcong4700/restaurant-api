package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204)
	})

	r.GET("/ping", func(c *gin.Context) {
		// Trả về dữ liệu dạng JSON với mã HTTP Status 200 (OK)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	r.Run(":" + port)

}
