package main

import (
	"goweb_clase4_tm/cmd/server/handler"
	"goweb_clase4_tm/internal"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	repository := internal.NewRepository("transactions.json")
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
