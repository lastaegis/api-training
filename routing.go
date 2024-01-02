package main

import (
	"github.com/labstack/echo/v4"
	"latihan-api/module/provinsi"
)

func RoutingAPI(echo *echo.Echo) {
	// Routing Provinsi
	echo.GET("/provinsi", provinsi.GetProvinisiAll)
	echo.GET("/provinsi/:id", provinsi.GetProvinsiById)
	echo.POST("/provinsi", provinsi.InsertProvinsi)
	echo.PUT("/provinsi/:id", provinsi.UpdateProvinsi)
	echo.DELETE("/provinsi/:id", provinsi.DeleteProvinsi)

	// Start Server
	echo.Logger.Fatal(echo.Start(":8001"))
}
