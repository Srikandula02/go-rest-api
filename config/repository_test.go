package config

import (
	"database/sql"
	"go-api/models"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var u = &models.Employee{
	Id:           "12121",
	FirstName:    "Taylor",
	MiddleName:   "Kayla",
	LastName:     "Swift",
	Gender:       "Female",
	Salary:       995585,
	DOB:          time.Now(),
	Email:        "fgfgfgf@gmail.com",
	Phone:        "444444444",
	State:        "nsw",
	Postcode:     55555,
	AddressLine1: "dfdfd",
	AddressLine2: "ind",
	TFN:          "555",
	SuperBalance: 5522,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetAllEmployees(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	log.Println("checking repo in testfind", repo)
	defer func() {
		repo.Close()
	}()
	query := "SELECT (.+) FROM employee"
	rows := sqlmock.NewRows([]string{"Id", "firstname", "middlename", "lastname", "gender", "salary", "dob", "email", "phone", "state", "postcode", "addressline1", "addressline2", "tfn", "superbalance"}).
		AddRow(u.Id, u.FirstName, u.MiddleName, u.LastName, u.Gender, u.Salary, "1994-01-01", u.Email, u.Phone, u.State, u.Postcode, u.AddressLine1, u.AddressLine2, u.TFN, u.SuperBalance)

	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := GetEmployeesQuery(query, repo)
	log.Println("checking users from::::", users)
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
}

func TestGetEmpById(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	log.Print("db value in testfindbyid::::::", repo)
	defer func() {
		repo.Close()
	}()
	query := "SELECT * FROM employee WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"Id", "firstname", "middlename", "lastname", "gender", "salary", "dob", "email", "phone", "state", "postcode", "addressline1", "addressline2", "tfn", "superbalance"}).
		AddRow(u.Id, u.FirstName, u.MiddleName, u.LastName, u.Gender, u.Salary, "1994-01-01", u.Email, u.Phone, u.State, u.Postcode, u.AddressLine1, u.AddressLine2, u.TFN, u.SuperBalance)

	log.Println("printing rows::::", rows)
	mock.ExpectQuery(query).WithArgs(u.Id).WillReturnRows(rows)
	log.Println("checking the u.Id value", u.Id)

	user, err := ReturnRowById(query, u.Id, repo)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestCreateEmp(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO employee"

	//query := "INSERT INTO users \\(Id,firstname ,middlename ,lastname ,gender ,salary ,dob ,email , phone , state ,postcode, addressline1 ,addressline2, tfn, superbalance\\) VALUES \\(\\?, \\?, \\?, \\?,\\?, \\?, \\?, \\?,\\?, \\?, \\?, \\?,\\?, \\?, \\? \\)"

	mock.MatchExpectationsInOrder(false)
	mock.ExpectExec(query).WithArgs().WillReturnResult(sqlmock.NewResult(0, 1))
	//prep := mock.ExpectPrepare(query)
	//prep.ExpectExec().WithArgs(u.Id, u.FirstName, u.MiddleName, u.LastName, u.Gender, u.Salary, u.DOB, u.Email, u.Phone, u.State, u.Postcode, u.AddressLine1, u.AddressLine2, u.TFN, u.SuperBalance).WillReturnResult(sqlmock.NewResult(0, 1))
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	err := InsertEmployee(c, repo)
	log.Println("checking the creating test error", err)
	assert.NoError(t, err)
}

func TestCreateEmpError(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO employee \\(id,first_name ,middle_name ,last_name ,gender ,salary ,dob ,email , phone , state ,postcode, address_line1 ,address_line2, TFN, super_balance\\) VALUES \\(\\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Id, u.FirstName, u.MiddleName, u.LastName, u.Gender, u.Salary, u.DOB, u.Email, u.Phone, u.State, u.Postcode, u.AddressLine1, u.AddressLine2, u.TFN, u.SuperBalance).WillReturnResult(sqlmock.NewResult(0, 0))

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	err := InsertEmployee(c, repo)
	assert.Error(t, err)
}

func TestUpdateEmp(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE employee WHERE id = \\?"
	//query := "UPDATE employee SET first_name = \\?, middle_name = \\?, last_name = \\?,gender = \\?, salary = \\?, dob = \\?, email = \\?, phone = \\?, state = \\?,postcode = \\?, address_line1 = \\?, address_line2 = \\?, TFN = \\?, super_balance = \\?  WHERE id = \\?"

	mock.MatchExpectationsInOrder(false)
	mock.ExpectExec(query).WithArgs().WillReturnResult(sqlmock.NewResult(0, 1))

	//prep := mock.ExpectPrepare(query)
	//prep.ExpectExec().WithArgs(u.FirstName, u.MiddleName, u.LastName, u.Gender, u.Salary, u.DOB, u.Email, u.Phone, u.State, u.Postcode, u.AddressLine1, u.AddressLine2, u.TFN, u.SuperBalance, u.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	err := UpdateQuery(query, repo, u.Id, c)
	assert.NoError(t, err)
}

func TestDeleteEmp(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "DELETE FROM employee where id = ?"

	mock.MatchExpectationsInOrder(false)
	mock.ExpectExec(query).WithArgs().WillReturnResult(sqlmock.NewResult(0, 1))
	//prep := mock.ExpectPrepare(query)
	//prep.ExpectExec().WithArgs(u.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := Deletequery(repo, u.Id)
	log.Println("checing tthe error in delete", err)
	assert.NoError(t, err)
}
