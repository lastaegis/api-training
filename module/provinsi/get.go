package provinsi

import (
	"github.com/labstack/echo/v4"
	general_structure "latihan-api/module/general-structure"
	"latihan-api/mysql"
	"log"
	"net/http"
)

func GetProvinisiAll(c echo.Context) error {
	provinsi := c.QueryParam("provinsi")

	// Connection handling
	db := mysql.DBConnection()
	defer db.Close()

	listProvinsi := []Provinsi{}
	if provinsi != "" {
		err := db.Select(&listProvinsi, "SELECT ID, PROVINSI FROM PROVINSI WHERE PROVINSI LIKE '%"+provinsi+"%' AND DELETED_AT IS NULL")
		if err != nil {
			log.Panic(err)
		}
	} else {
		err := db.Select(&listProvinsi, "SELECT ID, PROVINSI FROM PROVINSI WHERE DELETED_AT IS NULL")
		if err != nil {
			log.Panic(err)
		}
	}

	result := &general_structure.ResponseGet{
		Status:    200,
		Message:   "success",
		TotalData: int32(len(listProvinsi)),
		Data:      listProvinsi,
	}
	return c.JSON(http.StatusOK, result)
}

func GetProvinsiById(c echo.Context) error {
	id := c.Param("id")

	db := mysql.DBConnection()
	defer db.Close()

	provinsi := Provinsi{}
	err := db.Get(&provinsi, "SELECT ID, PROVINSI FROM PROVINSI WHERE ID = "+id+" AND DELETED_AT IS NULL")
	if err != nil {
		emptyResult := &general_structure.ResponseGet{
			Status:  204,
			Message: "ID Provinsi " + id + " tidak tersedia",
			Data:    make([]Provinsi, 0),
		}
		return c.JSON(http.StatusOK, emptyResult)
	}

	result := &general_structure.ResponseGet{
		Status:    200,
		Message:   "success",
		TotalData: 1,
		Data:      provinsi,
	}

	return c.JSON(http.StatusOK, result)
}
