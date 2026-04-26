package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"log"
)

type attribute struct {
	RecId       int
	Name        string
	Value       int
	Create_Date string
}

type asBody struct {
	CharacterId int
	Test string
}

func getAttributesAndSkills(c *gin.Context, db *sql.DB) {
	var body asBody
	if err := c.BindJSON(&body); err != nil {
		 c.JSON(400, gin.H{"error": err.Error()})
        return
	}
	log.Print("body",body.CharacterId)
	attributesQuery := `
		select RecId, Name, Value, Create_Date
		from Attributes
		Where CharacterId = ?
	`
	rows, err := db.Query(attributesQuery,body.CharacterId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	defer rows.Close()

	var attributes []attribute

	for rows.Next() {
		var i attribute
		rowErr := rows.Scan(&i.RecId, &i.Name, &i.Value, &i.Create_Date)
		if rowErr != nil {
			c.IndentedJSON(http.StatusBadRequest, rowErr)
		}
		attributes = append(attributes, i)
	}

	c.IndentedJSON(http.StatusOK, attributes)
}