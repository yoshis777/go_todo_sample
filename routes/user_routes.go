package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserSignUp(ctx *gin.Context) {
	println("post/signup")

	username := ctx.PostForm("username")
	emailaddress := ctx.PostForm("emailaddress")
	password := ctx.PostForm("password")
	passwordConf := ctx.PostForm("passwordconfirmation")

	println("username:" + username)
	println("emailaddress" + emailaddress)
	println("password" + password)
	println("passwordConf" + passwordConf)

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080")
}

func UserLogin(ctx *gin.Context) {
	println("post/login")

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	println(username)
	println(password)

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080")
}
