package handler

import (
	service2 "github.com/PeterlitsZo/CHOOSE/back/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListAnswerImpl struct {
	*gin.Context
}

func (g *ListAnswerImpl) Handle() {
	db := service2.Db.Session(&gorm.Session{NewDB: true})
	if db.Error != nil {
		panic(db.Error)
	}
	db = db.Model(&service2.Answer{})
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	userName := g.Query("user_name")
	db.Select("question, answer").Joins("join questions on answers.question_id = questions.id where answers.user_name = ?", userName)
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	ret := make([]questionAndAnswer, 0)
	db.Find(&ret)
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	g.JSON(200, gin.H{
		"answerList": ret,
	})
}

type questionAndAnswer struct {
	Answer   string `json:"answer"`
	Question string `json:"question"`
}
