package handler

import "github.com/gin-gonic/gin"

func ping(g *gin.Context) {
	g.JSON(200, gin.H{
		"ping": "pong",
	})
}

func responseOK(g *gin.Context) {
	g.JSON(200, gin.H{})
}

func responseError(g *gin.Context, err error) {
	g.JSON(400, gin.H{
		"error": err.Error(),
	})
}
