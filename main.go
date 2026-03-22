package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type sheetItem struct {
	RecId       int
	Name        string
	Type        string
	Create_Date string
}

var sheetItems = []sheetItem{
	{RecId: 1, Name: "Test", Type: "test", Create_Date: ""},
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

func connectToMariaDB() (*sql.DB, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbLocation := os.Getenv("DB_LOCATION")
	database := os.Getenv("DATABASE")
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbLocation + ")/" + database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, dbErr := connectToMariaDB()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	fmt.Println("Running...")
	router := gin.Default()
	router.GET("/SheetList", func(c *gin.Context) { getSheetList(c, db) })

	router.Run("localhost:1213")
}
