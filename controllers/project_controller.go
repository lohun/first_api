package controllers

import (
	"net/http"
	"strconv"

	"github.com/first_api/models"
	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {

	projects, err := models.GetProjects()
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, projects)
}


func GetProject(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return
	}
	project, err := models.GetProject(id)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, project)
}
