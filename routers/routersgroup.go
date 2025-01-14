package routers

import "github.com/gin-gonic/gin"

func Group(r *gin.Engine) *gin.Engine {
	r.Static("/xxx", "./statics")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", JWTAuthMiddleware(), Home)
	r.GET("/login", LoginGet)
	r.POST("/login", LoginPost)
	r.GET("/register", RegisterGet)
	r.POST("/register", RegisterPost)
	//r.POST("/auth", authHandler)
	return r
}
