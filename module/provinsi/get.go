package provinsi

import (
	"github.com/labstack/echo/v4"
	"latihan-api/mysql"
	"log"
	"net/http"
)

type Provinsi struct {
	ID       int    `db:"ID"`
	PROVINSI string `db:"PROVINSI"`
}

func GetProvinisiAll(c echo.Context) error {
	provinsi := c.QueryParam("provinsi")

	// Connection handling
	db := mysql.DBConnection()
	defer db.Close()

	listProvinsi := []Provinsi{}
	if provinsi != "" {
		err := db.Select(&listProvinsi, "SELECT ID, PROVINSI FROM PROVINSI WHERE PROVINSI LIKE '%"+provinsi+"%'")
		if err != nil {
			log.Panic(err)
		}
	} else {
		err := db.Select(&listProvinsi, "SELECT ID, PROVINSI FROM PROVINSI")
		if err != nil {
			log.Panic(err)
		}
	}

	return c.JSON(http.StatusOK, listProvinsi)
}

func GetProvinsiById(c echo.Context) error {
	id := c.Param("id")

	db := mysql.DBConnection()
	defer db.Close()

	provinsi := Provinsi{}
	err := db.Get(&provinsi, "SELECT ID, PROVINSI FROM PROVINSI WHERE ID = "+id)
	if err != nil {
		log.Panic(err)
	}

	return c.JSON(http.StatusOK, provinsi)
}
