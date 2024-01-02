package provinsi

import (
	"github.com/labstack/echo/v4"
	general_structure "latihan-api/module/general-structure"
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
	_, err := transaction.NamedExec("UPDATE PROVINSI SET DELETED_AT = :timestamp, DELETED_BY = :executor, SYNC_STATUS = 0 WHERE ID = :id", map[string]interface{}{
		"id":        id,
		"timestamp": time.Now(),
		"executor":  executor,
	})

	if err != nil {
		log.Panic(err)
	}

	transaction.Commit()

	result := &general_structure.ResponseDelete{
		Status:  200,
		Message: "Data successfully delete",
	}

	return c.JSON(http.StatusOK, result)
}
