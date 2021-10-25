package config

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Close() {
	r.DB.Close()
}

//Get All Employees
func (r *Repository) GetEmployees(c *gin.Context) {
	log.Println("hitting get all employees method")
	query := "SELECT * FROM employee"
	rows, err := GetEmployeesQuery(query, r)
	CheckErr(err, c)
	GetResponse(rows, c)
}

//Get Employee by Id
func (r *Repository) GetEmployeeById(c *gin.Context) {
	log.Println("hitting get employee by Id method")
	id := c.Param("id")
	query := "SELECT * FROM employee WHERE id = ?"
	row, err := ReturnRowById(query, id, r)
	ReturnResponse(row, c)
	CheckErr(err, c)
}

//Add an employee
func (r *Repository) AddEmployee(c *gin.Context) {
	log.Println("hitting Add employee")
	err := InsertEmployee(c, r)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "action unsuccessful"})
	} else {
		c.JSON(http.StatusCreated, "Employee has been added successfully")
	}
}

//Delete an employee
func (r *Repository) DeleteEmployee(c *gin.Context) {
	log.Println("hitting delete employee")
	id := c.Param("id")
	err := Deletequery(r, id)
	CheckError(err, c)
}

//Update an employee
func (r *Repository) UpdateEmployee(c *gin.Context) {
	log.Println("hitting update employee")
	id := c.Param("id")
	query := "UPDATE employee SET first_name=?,middle_name=?,last_name=?,gender=?,salary=?,dob=?,email=?,phone=?,state=?,postcode=?,address_line1=?,address_line2=?,TFN=?,super_balance=? where id = ?"
	err := UpdateQuery(query, r, id, c)
	CheckError(err, c)
}
