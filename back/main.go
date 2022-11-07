package main

import (
	"github.com/PeterlitsZo/CHOOSE/back/service"
	"github.com/gin-gonic/gin"
)

func init() {
	service.Init()
}

func main() {
	r := gin.Default()

	registerRouter(r)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
