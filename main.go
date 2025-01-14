package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sxxpqp/go-github-actions/routers"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())
	r = routers.Group(r)
	r.Run(":8080")
}
