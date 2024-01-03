package kota_kabupaten

import (
	"github.com/labstack/echo/v4"
	general_structure "latihan-api/module/general-structure"
	"latihan-api/mysql"
	"log"
	"net/http"
	"time"
)

func InsertKotaKabupaten(c echo.Context) error {
	idProvinsi := c.FormValue("id_provinsi")
	kotaKabupaten := c.FormValue("kota_kabupaten")
	executor := c.FormValue("executor")

	db := mysql.DBConnection()
	defer db.Close()

	transaction := db.MustBegin()

	query := `INSERT INTO KOTA_KABUPATEN (ID_PROVINSI, KOTA_KABUPATEN, CREATED_AT, CREATED_BY, UPDATED_AT, UPDATED_BY, SYNC_STATUS) 
    VALUE (:idProvinsi, :kotaKabupaten, :createdAt, :createdBy, :updatedAt, :updatedBy, :syncStatus)`
	transaction.NamedExec(query, map[string]interface{}{
		"idProvinsi":    idProvinsi,
		"kotaKabupaten": kotaKabupaten,
		"createdAt":     time.Now(),
		"createdBy":     executor,
		"updatedAt":     time.Now(),
		"updatedBy":     executor,
		"syncStatus":    0,
	})

	errCommit := transaction.Commit()
	if errCommit != nil {
		log.Panic(errCommit)
	}

	result := &general_structure.ResponsePost{Status: http.StatusOK, Message: "Data successfully saved"}

	return c.JSON(http.StatusOK, result)
}