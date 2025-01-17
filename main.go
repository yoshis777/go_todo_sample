package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go_todo_sample/helpers/sessions"
	"github.com/go_todo_sample/routes"
)

func main() {
	router := gin.Default()

	//静的ファイルの事前ロード
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	// セッションストアの初期化
	store := sessions.NewDummyStore()
	router.Use(sessions.StartDefaultSession(store))

	// ルーティング
	router.GET("/", routes.Home)
	router.GET("/signup", routes.SignUp)
	router.GET("/login", routes.Login)

	// 「（ホスト）/user/*」
	user := router.Group("/user")
	{
		user.POST("/signup", routes.UserSignUp)
		user.POST("/login", routes.UserLogin)
		user.POST("/logout", routes.UserLogOut)
	}

	router.NoRoute(routes.NoRoute)

	// サーバ起動（macのネットワーク受信設定の確認回避）
	router.Run("127.0.0.1:8080")
}
