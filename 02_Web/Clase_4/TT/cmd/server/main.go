package main

import (
	"goweb_clase4_tt/cmd/server/handler"
	"goweb_clase4_tt/cmd/server/middleware"
	"goweb_clase4_tt/docs"
	"goweb_clase4_tt/internal"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Transactions API
// @version 1.0
// @description This is a sample Transactions server.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	repository := internal.NewRepository("transactions.json")
	service := internal.NewService(repository)
	handler := handler.NewHandler(service)

	docs.SwaggerInfo.Title = "Transactions API"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	trxGroup := r.Group("/transactions")
	trxGroup.Use(middleware.TokenMiddleware())
	trxGroup.GET("/:id", handler.GetById())
	trxGroup.GET("", handler.GetAll())
	trxGroup.POST("", handler.Create())
	trxGroup.PUT("/:id", handler.Update())
	trxGroup.PATCH("/:id", handler.UpdateCodigoAndMonto())
	trxGroup.DELETE("/:id", handler.Delete())

	r.Run()
}
