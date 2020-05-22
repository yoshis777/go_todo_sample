package routes

// ほぼcontrollerに近そう

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 先頭大文字で外部パッケージから可視
func Home(ctx *gin.Context) {
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
