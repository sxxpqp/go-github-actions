package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sxxpqp/go-github-actions/tools"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func RegisterGet(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{
		"login": "login",
	})
}
func RegisterPost(c *gin.Context) {
	// Retrieve form data
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	rePassword := c.PostForm("password")
	log.Println(name, email, password, rePassword)
	// Basic Validation
	if name == "" || email == "" || password == "" || rePassword == "" {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "All fields are required.",
			"login": "login",
		})
		return
	}
	// Check if the passwords match
	if password != rePassword {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "Passwords do not match.",
			"login": "login",
		})
		log.Println("密码不一致")
		return
	}

	//// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"error": "Internal server error, please try again.",
			"login": "login",
		})
		return
	}
	log.Println("通过密码验证了")
	// Create the user object
	user := tools.User{
		Name:     name,
		Email:    email,
		PassWord: string(hashedPassword),
	}
	log.Println("user:", user)

	// Save user to database
	if err := db.Create(&user).Error; err != nil {
		log.Println("Error saving user:", err)
		c.String(200, "用户或者email已注册")
		//time.Sleep(30)
		c.Redirect(http.StatusFound, "/register")
		return
	}
	log.Println("插入数据成功")
	// Redirect to login page after successful registration
	c.Redirect(http.StatusFound, "/login")
}
