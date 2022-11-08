package handler

import (
	"fmt"
	"github.com/PeterlitsZo/CHOOSE/back/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateAnswerImpl struct {
	*gin.Context
}

func (g *CreateAnswerImpl) Handle() {
	db := service.Db.Session(&gorm.Session{NewDB: true})
	if db.Error != nil {
		panic(db.Error)
	}
	db = db.Model(&service.Answer{})
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	var ans service.Answer
	if err := g.ShouldBind(&ans); err != nil {
		responseError(g.Context, err)
		return
	}
	fmt.Println(ans)
	db = db.Create(&ans)
	if db.Error != nil {
		responseError(g.Context, db.Error)
		return
	}
	ping(g.Context)
}
