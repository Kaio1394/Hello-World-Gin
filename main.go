package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	//Path of bootstrap
	r.Static("/vendor", "./static/vendor")
	//HTML path
	r.LoadHTMLGlob("templates/**/**")
	//Route
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "view/index.html", gin.H{
			"title": "Hello GIN",
		})
	})
	log.Println("Server started!")
	r.Run(":8081")
}
