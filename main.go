package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go_todo_sample/routes"
)

func main() {
	router := gin.Default()

	//静的ファイルの事前ロード
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	// ルーティング
	router.GET("/", routes.Home)
	router.GET("/signup", routes.SignUp)
	router.GET("/login", routes.Login)

	// 「（ホスト）/user/*」
	user := router.Group("/user")
	{
		user.POST("/signup", routes.UserSignUp)
	}

	router.NoRoute(routes.NoRoute)

	// サーバ起動
	router.Run(":8080")
}
