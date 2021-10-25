package config

import (
	"database/sql"
	"go-api/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func ReturnRowById(s string, id string, r *Repository) (*sql.Rows, error) {
	row, err := r.DB.Query(s, id)
	return row, err
}

func ReturnResponse(rows *sql.Rows, c *gin.Context) {
	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(&emp.Id, &emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.Gender, &emp.Salary, &emp.DOB, &emp.Email, &emp.Phone,
			&emp.State, &emp.Postcode, &emp.AddressLine1, &emp.AddressLine2, &emp.TFN, &emp.SuperBalance)
		if err != nil {
			c.JSON(404, gin.H{"error": "no record for this particular Id"})
		}

		//Adding mountebank
		super, err := getSuper(emp.Id) //get super balance
		log.Println("super value from mountebank is:::::::::::", super)
		if err != nil {
			c.JSON(500, "problem in fetching super details from 3rd party")
		}
		if super == 0 {
			c.JSON(500, "failed retrived super")
		}

		emp.SuperBalance = super
		c.IndentedJSON(http.StatusOK, emp)
	}
}

func getSuper(id string) (float64, error) {
	mb_url := strings.Replace("http://localhost:6000/ato/employee/?/super", "?", id, 1)
	resp, err := http.Get(mb_url)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()
	//converting response to byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in ReadAll method", err)
		return 0.0, err
	}
	//converting String to float, since super is float type
	superMb, err := strconv.ParseFloat(string(body), 64)
	log.Println("\033[31m", "super value after parsing string to float:::::", superMb)
	if err != nil {
		log.Println("Error in ParseFloat", err)
		return 0.0, err
	}
	return superMb, nil
}

func GetEmployeesQuery(s string, r *Repository) (*sql.Rows, error) {
	rows, err := r.DB.Query(s)
	return rows, err
}

func GetResponse(rows *sql.Rows, c *gin.Context) {
	var emps []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(&emp.Id, &emp.FirstName, &emp.MiddleName, &emp.LastName, &emp.Gender, &emp.Salary, &emp.DOB, &emp.Email, &emp.Phone, &emp.State, &emp.Postcode, &emp.AddressLine1, &emp.AddressLine2, &emp.TFN, &emp.SuperBalance); err != nil {
			c.JSON(404, gin.H{"error": "data not found"})
			log.Fatal(err)
		}
		emps = append(emps, emp)

		log.Println("emps object::::", emps)

		//looping around
		// for i := 0; i < len(emps)-1; i++ {
		// 	index := i
		// 	//wg.Add(1)
		// 	func(index int) {
		// 		log.Println("inside for loop")
		// 		//defer wg.Done()
		// 		super, err := getSuper(emps[index].Id)
		// 		log.Println("\033[31m", "super value:::in getAll:::", super)
		// 		if err != nil {
		// 			emps[index].SuperBalance = 0
		// 		}
		// 		emps[index].SuperBalance = super
		// 	}(index)
		// }

		//adding waitgroups
		var wg sync.WaitGroup
		for i := 0; i < len(emps)-1; i++ {
			index := i
			wg.Add(1)
			go func(index int) {
				log.Println("inside go routine")
				defer wg.Done()
				super, err := getSuper(emps[index].Id)
				log.Println("\033[31m", "super value:::in getAll:::", super)
				if err != nil {
					emps[index].SuperBalance = 0
				}
				emps[index].SuperBalance = super
			}(index)
		}
		wg.Wait()

		c.IndentedJSON(http.StatusOK, emp)
	}

}

func InsertEmployee(c *gin.Context, r *Repository) error {
	var emp models.Employee
	c.BindJSON(&emp)
	_, err := r.DB.Exec("INSERT INTO employee (id,first_name ,middle_name ,last_name ,gender ,salary ,dob ,email , phone , state ,postcode, address_line1 ,address_line2, TFN, super_balance) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", emp.Id, emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB, emp.Email, emp.Phone, emp.State, emp.Postcode, emp.AddressLine1, emp.AddressLine2, emp.TFN, emp.SuperBalance)
	return err
}

func Deletequery(r *Repository, id string) error {
	_, err := r.DB.Query("DELETE FROM employee WHERE id = ?", id)
	return err
}

func UpdateQuery(s string, r *Repository, id string, c *gin.Context) error {
	var emp models.Employee
	c.BindJSON(&emp)
	_, err := r.DB.Exec(s, emp.FirstName, emp.MiddleName, emp.LastName, emp.Gender, emp.Salary, emp.DOB, emp.Email, emp.Phone, emp.State, emp.Postcode, emp.AddressLine1, emp.AddressLine2, emp.TFN, emp.SuperBalance, id)
	return err
}

func CheckErr(err error, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		log.Fatal(err)
	}
}

func CheckError(err error, c *gin.Context) {
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "action unsuccessful"})
	} else {
		c.JSON(http.StatusOK, "Action Successful")
	}
}
