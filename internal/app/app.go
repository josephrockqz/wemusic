package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	fmt.Println("app init")

	router := gin.Default()
	router.GET("/albums", services.getAccessToken)
	router.Run("localhost:8080")
}
