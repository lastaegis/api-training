package kota_kabupaten

import (
	"fmt"
	"github.com/labstack/echo/v4"
	general_structure "latihan-api/module/general-structure"
	"latihan-api/mysql"
	"log"
	"net/http"
	"time"
)

func UpdateKotaKabupaten(c echo.Context) error {
	idKotaKabupaten := c.Param("id")
	idProvinsi := c.FormValue("id_provinsi")
	kotaKabupaten := c.FormValue("kota_kabupaten")
	executor := c.FormValue("executor")

	db := mysql.DBConnection()
	defer db.Close()

	transaction := db.MustBegin()

	// SQL Builder
	sql := `UPDATE KOTA_KABUPATEN SET`
	if idProvinsi != "" {
		sql += ` ID_PROVINSI = :idProvinsi, `
	} else {

	}
	if kotaKabupaten != "" {
		sql += ` KOTA_KABUPATEN = :kotaKabupaten, `
	}
	if executor == "" {
		fmt.Println("Test")
		result := &general_structure.ResponsePut{
			Status:  http.StatusUnprocessableEntity,
			Message: "Parameters executor missing",
		}

		return c.JSON(http.StatusUnprocessableEntity, result)
	}

	sql += `UPDATED_AT = :updatedAt, UPDATED_BY = :updatedBy WHERE ID = :idKotaKabupaten`

	_, err := transaction.NamedExec(sql, map[string]interface{}{
		"idKotaKabupaten": idKotaKabupaten,
		"idProvinsi":      idProvinsi,
		"kotaKabupaten":   kotaKabupaten,
		"updatedAt":       time.Now(),
		"updatedBy":       executor,
	})

	if err != nil {
		log.Panic(err)
	}

	transaction.Commit()

	result := general_structure.ResponsePut{
		Status:  http.StatusOK,
		Message: "Data successfully update",
	}

	return c.JSON(http.StatusOK, result)
}
