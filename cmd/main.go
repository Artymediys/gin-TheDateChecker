package main

import (
	"gin_TheDateChecker/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.Use(internal.IsPing())

	route.GET("/when/:year", internal.GetDays)
	// route.GET("/when/:year", internal.GetDays, internal.isPing())

	route.Run()
}
