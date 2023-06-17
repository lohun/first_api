package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/first_api/models"
	"github.com/gin-gonic/gin"
)

type TempProject struct {
	ProjectName        string `json:"projectName"`
	ProjectDescription string `json:"projectDescription"`
	StartDate          string `json:"startDate"`
	EndDate            string `json:"endDate"`
}

type Status struct {
	User   int64  `json:"user"`
	Status string `json:"status"`
}

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

func AddProject(c *gin.Context) {
	var project TempProject

	err := c.BindJSON(&project)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	dateOutlet := strings.Split(project.StartDate, "/")
	year, _ := strconv.ParseInt(dateOutlet[2], 10, 0)
	month, _ := strconv.ParseInt(dateOutlet[1], 10, 0)
	day, _ := strconv.ParseInt(dateOutlet[0], 10, 0)
	start := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.Local)
	secondDateOutlet := strings.Split(project.EndDate, "/")
	secondYear, _ := strconv.ParseInt(secondDateOutlet[2], 10, 0)
	secondMonth, _ := strconv.ParseInt(secondDateOutlet[1], 10, 0)
	secondDay, _ := strconv.ParseInt(secondDateOutlet[0], 10, 0)
	second := time.Date(int(secondYear), time.Month(secondMonth), int(secondDay), 0, 0, 0, 0, time.Local)

	if models.AddProject(project.ProjectName, project.ProjectDescription, start.Format("2006-1-2"), second.Format("2006-1-2")) {
		c.Status(http.StatusCreated)
		return
	} else {
		c.Status(http.StatusBadRequest)
	}
}

func UpdateProjectStatus(c *gin.Context) {
	var status Status

	c.BindJSON(&status)
	if models.UpdateProjectStatus(status.User, status.Status) {
		c.Status(http.StatusCreated)
		return
	}
	c.Status(http.StatusBadRequest)
}
