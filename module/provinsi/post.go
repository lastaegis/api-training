package provinsi

import (
	"github.com/labstack/echo/v4"
	general_structure "latihan-api/module/general_structure"
	"latihan-api/mysql"
	"log"
	"net/http"
	"time"
)

func InsertProvinsi(c echo.Context) error {
	provinsi := c.FormValue("provinsi")
	executor := c.FormValue("executor")

	db := mysql.DBConnection()
	defer db.Close()

	transaction := db.MustBegin()
	_, err := transaction.NamedExec(`INSERT INTO PROVINSI (PROVINSI, CREATED_AT, CREATED_BY, UPDATED_AT, UPDATED_BY, SYNC_STATUS) VALUE (:provinsi, :created_at, :created_by, :updated_at, :updated_by, :sync_status)`,
		map[string]interface{}{
			"provinsi":    provinsi,
			"created_at":  time.Now(),
			"created_by":  executor,
			"updated_at":  time.Now(),
			"updated_by":  executor,
			"sync_status": 0,
		})
	if err != nil {
		log.Panic(err)
	}
	transaction.Commit()

	result := &general_structure.ResponsePost{
		Status:  http.StatusOK,
		Message: "Data provinsi successfully save",
	}

	return c.JSON(http.StatusOK, result)
}
