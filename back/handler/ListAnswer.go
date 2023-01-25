package handler

import (
	"github.com/PeterlitsZo/CHOOSE/back/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListAnswerImpl struct {
	*gin.Context
}

func (g *ListAnswerImpl) Handle() {
	db := service.Db.Session(&gorm.Session{NewDB: true})
	if db.Error != nil {
		panic(db.Error)
	}
	db = db.Model(&service.Answer{})
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	userName := g.Query("user_name")
	if userName != "" {
		db = db.Where("user_name = ?", userName)
	}
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	ret := make([]service.Answer, 0)
	db.Find(&ret)
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	g.JSON(200, gin.H{
		"answers": ret,
	})
}
