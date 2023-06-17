package controllers

import (
	"net/http"
	"strconv"

	"github.com/first_api/models"
	"github.com/gin-gonic/gin"
)

type TempTask struct {
	taskName        string
	taskDescription string
	startDate       string
	endDate         string
	assignedTo      int64
	project         int64
}

type TaskStatus struct {
	id   int64
	status string
}

func GetTasks(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return
	} 
	tasks, err := models.GetTasks(id)

	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return
	}
	task, err := models.GetTask(id)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func AddTask(c *gin.Context) {
	var task TempTask

	err := c.BindJSON(&task)
	if err != nil {
		return
	}

	if models.AddTask(task.taskName, task.taskDescription, task.startDate, task.endDate, task.assignedTo, task.project) {
		c.Status(http.StatusCreated)
		return
	}
	c.Status(http.StatusBadRequest)
}

func UpdateTaskStatus(c *gin.Context) {
	var status TaskStatus

	c.BindJSON(&status)
	if models.UpdateStatus(status.id, status.status) {
		c.Status(http.StatusCreated)
		return
	}
	c.Status(http.StatusBadRequest)
}
