package controllers

import (
	"net/http"
	"strconv"

	"github.com/first_api/models"
	"github.com/gin-gonic/gin"
)

type TempStakeHolder struct {
	name     string
	email    string
	phone    int64
	password string
	role     int64
	project  int64
}

type StakeHolderAsign struct {
	project     int64
	stakeHolder int64
	role        int64
}

func GetstakeHolders(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return
	}
	stakeHolders, err := models.GetStakeholders(id)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, stakeHolders)
}

func GetstakeHolder(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return
	}

	project := c.Param("id")
	projectId, err := strconv.ParseInt(project, 10, 64)
	if err != nil {
		return
	}
	stakeHolder, err := models.GetStakeholder(id, projectId)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, stakeHolder)
}

func AddstakeHolder(c *gin.Context) {
	var stakeHolder TempStakeHolder

	err := c.BindJSON(&stakeHolder)
	if err != nil {
		return
	}

	if models.AddStakeholder(stakeHolder.name, stakeHolder.email, stakeHolder.phone, stakeHolder.password, stakeHolder.role, stakeHolder.project) {
		c.Status(http.StatusCreated)
		return
	}
	c.Status(http.StatusBadRequest)
}

func AssignStakeHolder(c *gin.Context) {
	var asign StakeHolderAsign

	c.BindJSON(&asign)
	if models.AssignStakeHolder(asign.project, asign.stakeHolder, asign.role) {
		c.Status(http.StatusCreated)
		return
	}
	c.Status(http.StatusBadRequest)
}
