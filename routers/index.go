package routers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"login": "login",
	})
}
func Home(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{
		"login": "login",
	})
}
