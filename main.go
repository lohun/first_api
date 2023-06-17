package main

import (
	"github.com/gin-gonic/gin"
	"github.com/first_api/controllers"
)

func main() {
	router := gin.Default()
	// stakeholders
	router.POST("/stakeholer/add", controllers.AddstakeHolder)
	router.GET("/stakeholders/:id", controllers.GetstakeHolders)
	router.GET("/stakeholder/:id", controllers.GetstakeHolder)
	router.PUT("/stakeholder/assign", controllers.AssignStakeHolder)

	// tasks
	router.POST("/task/add", controllers.AddTask)
	router.GET("/task", controllers.GetTasks)
	router.GET("/task/:id", controllers.GetTask)
	router.PUT("/task/status", controllers.UpdateTaskStatus)

	// project
	router.POST("/project/add", controllers.AddProject)
	router.GET("/project", controllers.GetProjects)
	router.GET("/project/:id", controllers.GetProject)
	router.PUT("/project/status", controllers.UpdateProjectStatus)
	router.Run(":8080")
}
