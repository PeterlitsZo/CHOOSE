package main

import (
	"github.com/PeterlitsZo/CHOOSE/back/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	service.Init()
}

func main() {
	//gin.SetMode(gin.ReleaseMode) // 生产模式
	r := gin.Default()
	r.Use(cors.Default())

	registerRouter(r)

	if err := r.Run(":8033"); err != nil {
		panic(err)
	}
}
