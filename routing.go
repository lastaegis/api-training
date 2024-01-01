package main

import (
	"github.com/labstack/echo/v4"
	"latihan-api/lokasi"
)

func RoutingAPI(echo *echo.Echo) {
	// Routing Provinsi
	echo.GET("/provinsi/", lokasi.GetProvinisiAll)
	echo.GET("/provinsi/:id", lokasi.GetProvinsiById)

	// Start Server
	echo.Logger.Fatal(echo.Start(":8001"))
}
