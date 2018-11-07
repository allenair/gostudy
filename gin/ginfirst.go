package mygin

import (
	"fmt"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

// MainGin is entrance function
func MainGin() {
	fmt.Println("Hello!!!")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8000")
}
