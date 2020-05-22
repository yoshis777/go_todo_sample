package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/go_todo_sample/config"
)

func UserSignUp(ctx *gin.Context) {
	println("post/signup")

	username := ctx.PostForm("username")
	emailaddress := ctx.PostForm("emailaddress")
	password := ctx.PostForm("password")
	passwordConf := ctx.PostForm("passwordconfirmation")

	if password != passwordConf {
		println("Error: password and conf not match")
		ctx.Redirect(http.StatusSeeOther, "//localhost:8080")
		return
	}

	db := config.DummyDB()
	if err := db.SaveUser(username, emailaddress, password); err != nil {
		println("Error: " + err.Error())
	} else {
		println("username:" + username)
		println("emailaddress" + emailaddress)
		println("password" + password)
		println("passwordConf" + passwordConf)
	}

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080")
}

func UserLogin(ctx *gin.Context) {
	println("post/login")

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	db := config.DummyDB()
	user, err := db.GetUser(username, password)
	if err != nil {
		println("Error: " + err.Error())
	} else {
		println(username)
		println(password)

		user.Authenticate()
	}

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080")
}
