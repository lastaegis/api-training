package main

import (
	"github.com/labstack/echo/v4"
	kota_kabupaten "latihan-api/module/kota-kabupaten"
	"latihan-api/module/provinsi"
)

func RoutingAPI(echo *echo.Echo) {
	// Routing Provinsi
	echo.GET("/provinsi", provinsi.GetProvinisiAll)
	echo.GET("/provinsi/:id", provinsi.GetProvinsiById)
	echo.POST("/provinsi", provinsi.InsertProvinsi)
	echo.PUT("/provinsi/:id", provinsi.UpdateProvinsi)
	echo.DELETE("/provinsi/:id", provinsi.DeleteProvinsi)

	// Routing Kota/Kabupaten
	echo.GET("/kota-kabupaten", kota_kabupaten.GetKotaKabupatenAll)
	echo.GET("/kota-kabupaten/:id", kota_kabupaten.GetKotaKabupatenById)
	echo.POST("/kota-kabupaten", kota_kabupaten.InsertKotaKabupaten)
	echo.PUT("/kota-kabupaten/:id", kota_kabupaten.UpdateKotaKabupaten)
	echo.DELETE("/kota-kabupaten/:id", kota_kabupaten.DeleteKotaKabupaten)

	// Start Server
	echo.Logger.Fatal(echo.Start(":8001"))
}
