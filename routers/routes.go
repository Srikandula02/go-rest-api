package routers

import (
	"go-api/config"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	var data config.Repository
	DB := config.GetDB()
	data = config.Repository{DB}
	router := gin.Default()
	log.Println("Before the end points in SetupRouter function")

	router.GET("/employees", func(c *gin.Context) { data.GetEmployees(c) })
	router.GET("/employee/:id", func(c *gin.Context) { data.GetEmployeeById(c) })
	router.POST("/employee", func(c *gin.Context) { data.AddEmployee(c) })
	router.DELETE("/employee/:id", func(c *gin.Context) { data.DeleteEmployee(c) })
	router.PUT("/employee/:id", func(c *gin.Context) { data.UpdateEmployee(c) })
	router.Run("localhost:8080")
}
