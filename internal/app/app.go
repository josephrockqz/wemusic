package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/josephrockqz/wemusic-golang/internal/services"
)

func Run() {
	fmt.Println("app init")

	router := gin.Default()
	router.GET("/access-token", services.GetAccessToken)
	router.Run("localhost:3000")
}
