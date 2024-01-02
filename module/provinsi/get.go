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

	result := map[string]interface{}{
		"status":  200,
		"message": "success",
		"data":    listProvinsi,
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
		return c.NoContent(http.StatusNoContent)
	}

	result := map[string]interface{}{
		"status":  200,
		"message": "success",
		"data":    provinsi,
	}

	return c.JSON(http.StatusOK, result)
}
