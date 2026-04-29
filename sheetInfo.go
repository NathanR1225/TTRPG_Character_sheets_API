package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type attribute struct {
	RecId       int
	Name        string
	Value       int
	Create_Date string
	skills      []skill
}

type skill struct {
	RecId             int
	Name              string
	ParentAttributeId int
	InheritValue      bool
	Value             int
	CharacterId       int
}

type asBody struct {
	CharacterId int
	Test        string
}

func getAttributesAndSkills(c *gin.Context, db *sql.DB) {
	var body asBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Print("body", body.CharacterId)
	attributesQuery := `
		select RecId, Name, Value, Create_Date
		from Attributes
		Where CharacterId = ?
	`
	rows, err := db.Query(attributesQuery, body.CharacterId)
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
		i.skills = getSkills(i.RecId,db)
		attributes = append(attributes, i)
	}
	log.Print(" attributes",  attributes)
	
	c.JSON(http.StatusOK, gin.H{"attributes":attributes})
	log.Print(" done")
}

func getSkills(ParentAttributeId int, db *sql.DB) []skill {
	log.Print("ParentAttributeId", ParentAttributeId)
	skillQuery := `
		select RecId, Name, ParentAttributeId, InheritValue, COALESCE(Value,0), CharacterId
		from Skills
		Where ParentAttributeId = ?
	`
	rows, err := db.Query(skillQuery, ParentAttributeId)
	if err != nil {
		log.Panic("Unable to get skill for", ParentAttributeId)
		return nil
	}
	var skills []skill
	for rows.Next() {
		var s skill
		rowErr := rows.Scan(&s.RecId, &s.Name, &s.ParentAttributeId, &s.InheritValue, &s.Value, &s.CharacterId)
		if rowErr != nil {
			log.Panic("Error mapping skill  :", ParentAttributeId, "   ",rowErr)
		}
		skills = append(skills, s)
	}
	return skills
}
