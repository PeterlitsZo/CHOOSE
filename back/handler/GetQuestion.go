package handler

import (
	"github.com/PeterlitsZo/CHOOSE/back/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetQuestionImpl struct {
	*gin.Context
}

func (g *GetQuestionImpl) Handle() {
	db := service.Db.Session(&gorm.Session{NewDB: true})
	if db.Error != nil {
		panic(db.Error)
	}
	db = db.Model(&service.Question{})
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	id := g.Param("id")
	var ques service.Question
	db.Where("id = ?", id).Scan(&ques)
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	g.JSON(200, ques)
}
