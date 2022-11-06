package main

import (
	"github.com/PeterlitsZo/CHOOSE/back/handler"
	"github.com/gin-gonic/gin"
)

func CreateAnswer(g *gin.Context) {
	h := handler.CreateAnswerImpl{
		Context: g,
	}
	h.Handle()
}

func ListAnswer(g *gin.Context) {
	h := handler.ListAnswerImpl{
		Context: g,
	}
	h.Handle()
}

func GetQuestion(g *gin.Context) {
	h := handler.GetQuestionImpl{
		Context: g,
	}
	h.Handle()
}
