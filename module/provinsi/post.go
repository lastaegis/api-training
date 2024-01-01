package provinsi

import (
	"github.com/labstack/echo/v4"
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
	_, err := transaction.NamedExec(`INSERT INTO PROVINSI (PROVINSI, CREATED_AT, CREATED_BY, UPDATED_AT, UPDATED_BY) VALUE (:provinsi, :created_at, :created_by, :updated_at, :updated_by)`,
		map[string]interface{}{
			"provinsi":   provinsi,
			"created_at": time.Now(),
			"created_by": executor,
			"updated_at": time.Now(),
			"updated_by": executor,
		})
	if err != nil {
		log.Panic(err)
	}
	transaction.Commit()

	result := map[string]interface{}{
		"message": "Data saved to Database",
	}

	return c.JSON(http.StatusOK, result)
}
