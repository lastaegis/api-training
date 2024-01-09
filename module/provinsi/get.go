package provinsi

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"latihan-api/helper"
	"latihan-api/module/general_structure"
	"latihan-api/mysql"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

// Daftar query parameters yang disediakan pada endpoint: /provinsi?query_params
// Apabila terdapat penambahan parameter pencarian dapat dilakukan register
// pada data slice dibawah
var listAvailableQueryParams = []string{
	"provinsi",
	"limit",
}

var defaultLimit = 10

// GetProvinisiAll Logic yang digunakan untuk mendapatkan data seluruh provinsi
// ataupun melakukan pencarian berdasarkan paramaters yang disediakan
// Paramaters yang saat ini disediakan adalah:
// 1. provinsi <string>
func GetProvinisiAll(c echo.Context) error {
	// Validate all incoming params
	resultQueryParamValidation := validateQueryParams(c.QueryParams())

	if resultQueryParamValidation != nil {
		return c.JSON(http.StatusBadRequest, resultQueryParamValidation)
	}

	// Get Post Data
	provinsi := c.QueryParam("provinsi")
	limit := c.QueryParam("limit")

	// Connection handling
	db := mysql.DBConnection()
	defer db.Close()

	listProvinsi := []Provinsi{}
	sql := `SELECT ID, PROVINSI FROM PROVINSI `
	if provinsi != "" {
		sql = sql + `WHERE PROVINSI LIKE '%` + provinsi + `%' AND DELETED_AT IS NULL `
	} else {
		sql = sql + `WHERE DELETED_AT IS NULL `
	}

	if limit != "" {
		sql = sql + "LIMIT " + limit
	} else {
		sql = sql + "LIMIT " + strconv.FormatInt(int64(defaultLimit), defaultLimit)
	}

	fmt.Println(sql)

	err := db.Select(&listProvinsi, sql)
	if err != nil {
		log.Panic(err)
	}

	lenProvinsi := len(listProvinsi)
	if lenProvinsi == 0 {
		result := &general_structure.ResponseGet{
			Status:    http.StatusNoContent,
			Message:   "Data provinsi dengan paramaters yang diminta tidak tersedia",
			TotalData: int32(lenProvinsi),
			Data:      listProvinsi,
		}

		return c.JSON(http.StatusOK, result)
	}

	result := &general_structure.ResponseGet{
		Status:    http.StatusOK,
		Message:   "success",
		TotalData: int32(lenProvinsi),
		Data:      listProvinsi,
	}
	return c.JSON(http.StatusOK, result)
}

// GetProvinsiById logic yang digunakan untuk melakukan pengambilan
// data provinsi berdasarkan ID Provinsi
func GetProvinsiById(c echo.Context) error {
	id := c.Param("id")

	db := mysql.DBConnection()
	defer db.Close()

	provinsi := Provinsi{}
	err := db.Get(&provinsi, "SELECT ID, PROVINSI FROM PROVINSI WHERE ID = "+id+" AND DELETED_AT IS NULL")
	if err != nil {
		emptyResult := &general_structure.ResponseGet{
			Status:  http.StatusNoContent,
			Message: "ID Provinsi " + id + " tidak tersedia",
			Data:    make([]Provinsi, 0),
		}
		return c.JSON(http.StatusOK, emptyResult)
	}

	result := &general_structure.ResponseGet{
		Status:    http.StatusOK,
		Message:   "success",
		TotalData: 1,
		Data:      provinsi,
	}

	return c.JSON(http.StatusOK, result)
}

// Digunakan untuk melakukan validasi seluruh inbound query parameters
// untuk memastikan bahwa query parameter yang dikirim memang telah disediakan
// dan apabila didapatkan query parameter yang tidak terdaftar, maka akan dilakukan
// response kembali query parameter tidak ditemukan
func validateQueryParams(inboundQueryParams interface{}) interface{} {
	// Declare result var
	var result interface{}

	// Manipulate interface inbound query params for loop
	inboundQPReflect := reflect.ValueOf(inboundQueryParams)
	if inboundQPReflect.Len() > 0 {
		for _, key := range inboundQPReflect.MapKeys() {
			bv := helper.FindNeedle(listAvailableQueryParams, key.String())
			if !bv {
				result = *&general_structure.ResponseBadRequest{
					Message: key.String() + " tidak tersedia sebagai paramater",
				}
			}
		}
	}

	return result
}
