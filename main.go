package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go_todo_sample/routes"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	router.GET("/", routes.Home)
	router.NoRoute(routes.NoRoute)

	router.Run(":8080")
}
