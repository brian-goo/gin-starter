package db

import (
	"database/sql"
	"gtest/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetRows(c *gin.Context, host *string, query *string, params ...interface{}) (*sql.Rows, error) {
	var r *sql.Rows

	db, err := sql.Open("mysql", *host)
	if err != nil {
		handler.LogErr(c, "db error: "+err.Error())
		return r, err
	}
	// defer db.Close()

	rows, err := db.Query(*query, params...)
	if err != nil {
		handler.LogErr(c, "sql error: "+err.Error())
		return r, err
	}

	return rows, nil
}
