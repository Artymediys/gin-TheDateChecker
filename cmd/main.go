package main

import (
	"gin_TheDateChecker/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/when/:year", internal.GetDays)

	route.Run()
}
