package provinsi

import (
	"github.com/labstack/echo/v4"
	"latihan-api/mysql"
	"log"
	"net/http"
)

type Provinsi struct {
	ID       int32  `json:"id" db:"ID"`
	PROVINSI string `json:"provinsi" db:"PROVINSI"`
}

type Response struct {
	Status  int32       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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

	result := &Response{
		Status:  200,
		Message: "success",
		Data:    listProvinsi,
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
		emptyResult := &Response{
			Status:  204,
			Message: "ID Provinsi " + id + " tidak tersedia",
			Data:    make([]Provinsi, 0),
		}
		return c.JSON(http.StatusOK, emptyResult)
	}

	result := &Response{
		Status:  200,
		Message: "success",
		Data:    provinsi,
	}

	return c.JSON(http.StatusOK, result)
}
