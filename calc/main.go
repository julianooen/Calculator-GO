package main

import (
	"calc/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/calc/:operation/:num1/:num2", handlers.CalcHandler)
	router.GET("/calc/history", handlers.History)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
