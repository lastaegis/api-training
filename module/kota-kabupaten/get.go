package kota_kabupaten

import (
	"github.com/labstack/echo/v4"
	"latihan-api/mysql"
	"log"
	"net/http"
)

type StructKotaKabupaten struct {
	ID             int    `db:"ID"`
	PROVINSI       string `db:"PROVINSI"`
	KOTA_KABUPATEN string `db:"KOTA_KABUPATEN"`
}

func GetKotaKabupatenAll(c echo.Context) error {
	idProvinsi := c.QueryParam("id_provinsi")
	provinsi := c.QueryParam("provinsi")
	kotaKabupaten := c.QueryParam("kota_kabupaten")

	db := mysql.DBConnection()
	defer db.Close()

	listKotaKabupaten := []StructKotaKabupaten{}
	query := "SELECT KOTA_KABUPATEN.ID, PROVINSI, KOTA_KABUPATEN " +
		"FROM KOTA_KABUPATEN " +
		"INNER JOIN PROVINSI ON KOTA_KABUPATEN.ID_PROVINSI = PROVINSI.ID " +
		"WHERE KOTA_KABUPATEN.DELETED_AT IS NULL"
	if idProvinsi != "" {
		query += " AND ID_PROVINSI = " + idProvinsi
	}
	if provinsi != "" {
		query += " AND PROVINSI like '%" + provinsi + "%'"
	}
	if kotaKabupaten != "" {
		query += " AND KOTA_KABUPATEN like '%" + kotaKabupaten + "%'"
	}

	//Logging
	log.Println(query)

	err := db.Select(&listKotaKabupaten, query)

	if err != nil {
		log.Panic(err)
	}

	// Render result
	result := map[string]interface{}{
		"status_code": 200,
		"message":     "success",
		"data":        listKotaKabupaten,
	}

	return c.JSON(http.StatusOK, result)
}

func GetKotaKabupatenById(c echo.Context) error {
	id := c.Param("id")

	db := mysql.DBConnection()
	defer db.Close()

	dataKotaKabupaten := StructKotaKabupaten{}
	query := "SELECT KOTA_KABUPATEN.ID, PROVINSI, KOTA_KABUPATEN " +
		"FROM KOTA_KABUPATEN " +
		"INNER JOIN PROVINSI ON KOTA_KABUPATEN.ID_PROVINSI = PROVINSI.ID " +
		"WHERE KOTA_KABUPATEN.ID = " + id +
		" AND KOTA_KABUPATEN.DELETED_AT IS NULL"

	err := db.Get(&dataKotaKabupaten, query)
	if err != nil {
		log.Panic(err)
	}

	// Render result
	result := map[string]interface{}{
		"status_code": 200,
		"message":     "success",
		"data":        dataKotaKabupaten,
	}

	return c.JSON(http.StatusOK, result)
}
