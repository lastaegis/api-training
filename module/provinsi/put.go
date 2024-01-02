package provinsi

import (
	"github.com/labstack/echo/v4"
	"latihan-api/mysql"
	"log"
	"net/http"
	"time"
)

func UpdateProvinsi(c echo.Context) error {
	id := c.Param("id")
	provinsi := c.FormValue("provinsi")
	executor := c.FormValue("executor")

	// Open DB Connection
	db := mysql.DBConnection()
	defer db.Close()

	transaction := db.MustBegin()
	_, err := transaction.NamedExec("UPDATE PROVINSI SET PROVINSI = :provinsi, UPDATED_AT = :timestamp, UPDATED_BY = :executor, SYNC_STATUS = 0 WHERE ID = :id", map[string]interface{}{
		"id":        id,
		"provinsi":  provinsi,
		"executor":  executor,
		"timestamp": time.Now(),
	})

	if err != nil {
		log.Panic(err)
	}
	transaction.Commit()

	result := map[string]interface{}{
		"message": "Data provinsi berhasil diupdate",
	}

	return c.JSON(http.StatusOK, result)
}
