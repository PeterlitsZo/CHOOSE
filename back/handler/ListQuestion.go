package handler

import (
	"fmt"

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
		responseError(g.Context, fmt.Errorf("get the sesstion: %w", db.Error))
		return
	}
	db = db.Model(&service.Question{})
	if db.Error != nil {
		responseError(g.Context, fmt.Errorf("get question model: %w", db.Error))
		return
	}
	ret := make([]service.Question, 0)
	db.Preload("Choices").Find(&ret)
	if db.Error != nil {
		responseError(g.Context, fmt.Errorf("find all questions: %w", db.Error))
		return
	}
	g.JSON(200, gin.H{
		"questions": ret,
	})
}
