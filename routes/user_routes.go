package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go_todo_sample/helpers/sessions"
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
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	db := config.DummyDB()
	if err := db.SaveUser(username, emailaddress, password); err != nil {
		println("Error: " + err.Error())
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	println("Signup success!!")
	println("username:" + username)
	println("emailaddress" + emailaddress)
	println("password" + password)
	println("passwordConf" + passwordConf)

	user, err := db.GetUser(username, password)
	if err != nil {
		println("Error: while loading user: " + err.Error())
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	session := sessions.GetDefaultSession(ctx)
	session.Set("user", user)
	session.Save()
	println("Session saved.")
	println(" sessionID: " + session.ID)

	ctx.Redirect(http.StatusSeeOther, "/")
}

func UserLogin(ctx *gin.Context) {
	println("post/login")

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	db := config.DummyDB()
	user, err := db.GetUser(username, password)
	if err != nil {
		println("Error: " + err.Error())
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	println("Authentication Success!!")
	println(user.Username)
	println(user.Email)
	println(user.Password)

	session := sessions.GetDefaultSession(ctx)
	session.Set("user", user)
	session.Save()
	user.Authenticate()

	println("Session saved")
	println(" sessionID " + session.ID)
	ctx.Redirect(http.StatusSeeOther, "/")
}
