package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type sheetItem struct {
	RecId       int
	Name        string
	Type        string
	Create_Date string
}

func getSheetList(c *gin.Context, db *sql.DB) {
	query := "select RecId, Name, Type, Create_Date from Character_sheets_List"
	rows, err := db.Query(query)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	defer rows.Close()

	var sheetItems []sheetItem

	for rows.Next() {
		var i sheetItem
		rowErr := rows.Scan(&i.RecId, &i.Name, &i.Type, &i.Create_Date)
		if rowErr != nil {
			c.IndentedJSON(http.StatusBadRequest, rowErr)
		}
		sheetItems = append(sheetItems, i)
	}

	c.IndentedJSON(http.StatusOK, sheetItems)
}