package main

import (
	"goweb_clase3_tm/cmd/server/handler"
	"goweb_clase3_tm/internal"

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
	trxGroup.PUT("/:id", handler.Update())
	trxGroup.PATCH("/:id", handler.UpdateCodigoAndMonto())
	trxGroup.DELETE("/:id", handler.Delete())

	r.Run()
}
