package kota_kabupaten

import (
	"github.com/labstack/echo/v4"
	general_structure "latihan-api/module/general-structure"
	"latihan-api/mysql"
	"log"
	"net/http"
)

func DeleteKotaKabupaten(echo echo.Context) error {
	idKotaKabupaten := echo.Param("id")
	
	db := mysql.DBConnection()
	defer db.Close()
	
	transaction := db.MustBegin()
	query := `DELETE KOTA_KABUPATEN WHERE ID = :idKotaKabupaten`
	_, errQuery := transaction.NamedExec(query, map[string]interface{}{
		"idKotaKabupaten": idKotaKabupaten,
	})
	if errQuery != nil {
		log.Panic(errQuery)
	}
	
	errCommit := transaction.Commit()
	if errCommit != nil {
		log.Panic(errCommit)
	}
	
	result := &general_structure.ResponseDelete{
		Status: 200,
		Message: "Data successfully delete",
	}
	
	return echo.JSON(http.StatusOK, result)
}