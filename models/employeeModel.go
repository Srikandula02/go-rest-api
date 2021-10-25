package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id           string    `json:"id"`
	FirstName    string    `json:"firstname"`
	MiddleName   string    `json:"middlename"`
	LastName     string    `json:"lastname"`
	Gender       string    `json:"gender"`
	Salary       float64   `json:"salary"`
	DOB          time.Time `json:"dob"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	State        string    `json:"state"`
	Postcode     int       `json:"postcode"`
	AddressLine1 string    `json:"addressline1"`
	AddressLine2 string    `json:"addressline2"`
	TFN          string    `json:"tfn"`
	SuperBalance float64   `json:"superbalance"`
}

type RepositoryInterface interface {
	Close()
	GetEmployeeById(c *gin.Context)
	GetEmployees(c *gin.Context)
	AddEmployee(c *gin.Context)
	DeleteEmployee(c *gin.Context)
	UpdateEmployee(c *gin.Context)
}
