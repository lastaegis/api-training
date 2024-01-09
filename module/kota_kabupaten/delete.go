package kota_kabupaten

import (
	"github.com/labstack/echo/v4"
	general_structure "latihan-api/module/general_structure"
	"latihan-api/mysql"
	"log"
	"net/http"
	"time"
)

type ParametersDelete struct {
	IDKotaKabupaten string
	DeletedAt       time.Time
	DeletedBy       string
}

func DeleteKotaKabupaten(echo echo.Context) error {
	// Get Post Data
	idKotaKabupaten := echo.Param("id")
	executor := echo.FormValue("executor")

	// Add to structure
	parametersData := &ParametersDelete{
		IDKotaKabupaten: idKotaKabupaten,
		DeletedAt:       time.Now(),
		DeletedBy:       executor,
	}

	// Validate All Request Param
	validateResult := validateIdKotaKabupaten(parametersData)
	if validateResult != nil {
		return echo.JSON(http.StatusUnprocessableEntity, validateResult)
	}

	db := mysql.DBConnection()
	defer db.Close()

	transaction := db.MustBegin()
	query := `UPDATE KOTA_KABUPATEN SET DELETED_AT = :deletedAt, DELETED_BY = :deletedBy WHERE ID = :idKotaKabupaten`
	_, errQuery := transaction.NamedExec(query, map[string]interface{}{
		"idKotaKabupaten": idKotaKabupaten,
		"deletedAt":       time.Now(),
		"deletedBy":       executor,
	})
	if errQuery != nil {
		log.Panic(errQuery)
	}

	errCommit := transaction.Commit()
	if errCommit != nil {
		log.Panic(errCommit)
	}

	result := &general_structure.ResponseDelete{
		Status:  200,
		Message: "Data successfully delete",
	}

	return echo.JSON(http.StatusOK, result)
}

func validateIdKotaKabupaten(data *ParametersDelete) interface{} {
	validateData := []string{}

	if data.IDKotaKabupaten == "" {
		validateData = append(validateData, "ID Kota/Kabupaten Wajib Diisi")
	}

	result := &general_structure.ResponseFailed{
		Status:  http.StatusUnprocessableEntity,
		Message: validateData,
	}

	return result
}
