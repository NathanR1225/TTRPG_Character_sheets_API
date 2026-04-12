package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"time"
	// "net/http"
	"os"
)

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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//located in sheetList.go
	router.GET("/getSheetList", func(c *gin.Context) { getSheetList(c, db) })

	router.Run("localhost:1213")
}
