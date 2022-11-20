package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPerson godoc
// @Summary Get Person Information by Name
// @Description Get Person Information by Name
// @Tags Person
// @ID get-person-by-name
// @Accept  json
// @Produce  html
// @Param person_name path string true "Person Name"
// @Header 200 {string} Token "qwerty"
// @Router /persons/view/{person_name} [get]
func GetPerson(c *gin.Context) {
	if personName := c.Param("person_name"); personName != "" {
		// Check if the person exists
		if person, err := getPersonByName(personName); err == nil && person != nil {
			c.JSON(http.StatusOK, person)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// AddPerson godoc
// @Summary Add Person
// @Description Add Person By Information
// @Tags Person
// @ID add-person
// @Accept  json
// @Produce  json
// @Param   person body Person true "Person Information"
// @Success 200 {string} string
// @Header 200 {string} Token "qwerty"
// @Router /persons/add [post]
func AddPerson(c *gin.Context) {

	var person Person
	err := c.BindJSON(&person)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, err)
	}

	addNewPersonObject(person)

	c.JSON(http.StatusOK, "Person Added")

}
