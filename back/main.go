package main

import (
	"github.com/PeterlitsZo/CHOOSE/back/service"
	"github.com/gin-gonic/gin"
)

func init() {
	service.Init()
}

func main() {
	//gin.SetMode(gin.ReleaseMode) // 生产模式
	r := gin.Default()

	registerRouter(r)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
