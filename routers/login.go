package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sxxpqp/go-github-actions/tools"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var db = tools.Db()

func LoginGet(ctx *gin.Context) {
	ctx.HTML(200, "login.html", gin.H{
		"login": "login",
	})
}

func LoginPost(c *gin.Context) {
	//email := c.PostForm("email")
	//password := c.PostForm("password")
	//log.Println(email, password)
	//
	//var username tools.User
	//db.Where("Email = ?", email).First(&username)
	//log.Println([]byte(username.PassWord))
	//log.Println([]byte(password))
	//err := bcrypt.CompareHashAndPassword([]byte(username.PassWord), []byte(password))
	//if err != nil {
	//	c.Redirect(http.StatusFound, "/login")
	//	return
	//}
	//
	//c.Redirect(http.StatusFound, "/")
	email := c.PostForm("email")
	password := c.PostForm("password")
	log.Println(email, password)

	var user tools.User
	db.Where("Email = ?", email).First(&user)
	log.Println([]byte(user.PassWord))
	log.Println([]byte(password))
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenString, _ := GenToken(user.Name)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
