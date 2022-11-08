package handler

import (
	"github.com/PeterlitsZo/CHOOSE/back/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListQuestionImpl struct {
	*gin.Context
}

func (g *ListQuestionImpl) Handle() {
	db := service.Db.Session(&gorm.Session{NewDB: true})
	if db.Error != nil {
		panic(db.Error)
	}
	db = db.Model(&service.Question{})
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	ret := make([]service.Question, 0)
	db.Find(&ret)
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	g.JSON(200, gin.H{
		"questionList": ret,
	})
}
