package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Khởi tạo router mặc định của Gin (đã bao gồm Logger và Recovery middleware)
	r := gin.Default()

	// 2. Định nghĩa một Route đơn giản (GET request)
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
