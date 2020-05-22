package routes

// ほぼcontrollerに近そう

import (
	"github.com/gin-gonic/gin"
	"github.com/go_todo_sample/config"
	"github.com/go_todo_sample/helpers/sessions"

	"net/http"
)

// 先頭大文字で外部パッケージから可視
// Contextは状態をもつ、ポインタ渡しのため関数から永続的に変更可能
func Home(ctx *gin.Context) {
	// 宣言+代入
	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")

	if exists {
		user := buffer.(*config.DummyUserModel)
		println("user data is taken over")
		println("username: " + user.Username)
		println("password:" + user.Password)
	} else {
		println("user data discarded")
		println("sessionID" + session.ID)
	}

	session.Save()

	// ステータスコード、テンプレートファイル、ヘッダー定義
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}

func SignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", gin.H{})
}

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}
