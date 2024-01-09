package kota_kabupaten

import (
	"github.com/labstack/echo/v4"
	general_structure "latihan-api/module/general_structure"
	"latihan-api/mysql"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Variable Need To Sync mendandakan bahwa data
// Perlu dilakukan sinkronisasi ulang dengan
// data pada Datalake
var NeedToSync = 0

type ParamsInputKotaKabupaten struct {
	IDProvinsi    *int
	KotaKabupaten string
	CreatedAt     time.Time
	CreatedBy     string
	UpdatedAt     time.Time
	UpdatedBy     string
	SyncStatus    *int
}

// Digunakan untuk melakukan eksekusi insert data Kota/Kabupaten
func InsertKotaKabupaten(c echo.Context) error {
	idProvinsi, _ := strconv.Atoi(c.FormValue("id_provinsi"))
	kotaKabupaten := c.FormValue("kota_kabupaten")
	executor := c.FormValue("executor")

	dataFormValue := &ParamsInputKotaKabupaten{
		IDProvinsi:    &idProvinsi,
		KotaKabupaten: kotaKabupaten,
		CreatedAt:     time.Now(),
		CreatedBy:     executor,
		UpdatedAt:     time.Now(),
		UpdatedBy:     executor,
		SyncStatus:    &NeedToSync,
	}

	// Validasi seluruh parameters yang digunakan saat proses insert
	validationResult := validateRequiredParams(dataFormValue)
	if validationResult != nil {
		c.JSON(http.StatusBadRequest, validationResult)
	}

	db := mysql.DBConnection()
	defer db.Close()

	transaction := db.MustBegin()

	query := `INSERT INTO KOTA_KABUPATEN (ID_PROVINSI, KOTA_KABUPATEN, CREATED_AT, CREATED_BY, UPDATED_AT, UPDATED_BY, SYNC_STATUS) 
    VALUE (:idProvinsi, :kotaKabupaten, :createdAt, :createdBy, :updatedAt, :updatedBy, :syncStatus)`
	transaction.NamedExec(query, dataFormValue)

	errCommit := transaction.Commit()
	if errCommit != nil {
		log.Panic(errCommit)
	}

	result := &general_structure.ResponsePost{Status: http.StatusOK, Message: "Data successfully saved"}

	return c.JSON(http.StatusOK, result)
}

// Digunakan untuk melakukan validasi parameters yang dibutuhkan dalam
// Insert data Kota/Kabupaten ke Database
func validateRequiredParams(dataParams *ParamsInputKotaKabupaten) interface{} {
	validationMessage := []string{}

	// Validasi data ID Provinsi
	if dataParams.IDProvinsi == nil {
		validationMessage = append(validationMessage, "ID Provinsi merupakan isian wajib")
	}

	// Validasi data Kota/Kabupaten
	if dataParams.KotaKabupaten == "" {
		validationMessage = append(validationMessage, "Kota/Kabupaten merupakan isian wajib")
	}

	if dataParams.CreatedBy == "" {
		validationMessage = append(validationMessage, "Nama ekstekutor tidak terdeteksi dalam sistem")
	}

	return validationMessage
}
