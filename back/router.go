package main

import "github.com/gin-gonic/gin"

func registerRouter(r *gin.Engine) {
	registerChoose(r)
}

func registerChoose(r *gin.Engine) {
	choose := r.Group("/CHOOSE")
	choose.POST("/api/v1/answers", CreateAnswer)
	choose.GET("/api/v1/answers", ListAnswer)
	choose.GET("/api/v1/questions/:id", GetQuestion)
	choose.GET("/api/v1/questions", ListQuestion)
}
