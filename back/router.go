package main

import "github.com/gin-gonic/gin"

func registerRouter(r *gin.Engine) {
	registerChoose(r)
}

func registerChoose(r *gin.Engine) {
	r.POST("/api/v1/answers", CreateAnswer)
	r.GET("/api/v1/answers", ListAnswer)
	r.GET("/api/v1/questions/:id", GetQuestion)
	r.GET("/api/v1/questions", ListQuestion)
}
