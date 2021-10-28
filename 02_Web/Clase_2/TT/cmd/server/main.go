package main

import (
	"goweb_clase2_tt/cmd/server/handler"
	"goweb_clase2_tt/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repository := internal.NewRepository()
	service := internal.NewService(repository)
	handler := handler.NewHandler(service)

	trxGroup := r.Group("/transactions")
	trxGroup.GET("/:id", handler.GetById())
	trxGroup.GET("", handler.GetAll())
	trxGroup.POST("", handler.Create())

	r.Run()
}
