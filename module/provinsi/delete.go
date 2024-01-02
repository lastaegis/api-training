package provinsi

import (
	"github.com/labstack/echo/v4"
	"latihan-api/mysql"
	"log"
	"net/http"
	"time"
)

func DeleteProvinsi(c echo.Context) error {
	id := c.Param("id")
	executor := c.FormValue("executor")

	// Open DB Connection
	db := mysql.DBConnection()
	defer db.Close()

	transaction := db.MustBegin()
	_, err := transaction.NamedExec("UPDATE PROVINSI SET DELETE_AT = :timestamp, DELETED_BY = :executor, SYNC_STATUS = 0 WHERE ID = :id", map[string]interface{}{
		"id":        id,
		"timestamp": time.Now(),
		"executor":  executor,
	})

	if err != nil {
		log.Panic(err)
	}

	transaction.Commit()

	result := map[string]interface{}{
		"status":  200,
		"message": "Data delete success",
	}

	return c.JSON(http.StatusOK, result)
}
