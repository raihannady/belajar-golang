package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

type Employee struct {
	ID int
	Full_name string
	Email string
	Age int
	Division string
}


var (
	db *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connect to database!")

	// CreateEmployee()
	// GetEmployees()
	// UpdateEmployee()
	DeleteEmployee()
}

func CreateEmployee(){
	var employee = Employee{}
	sqlStatement := `INSERT INTO employees (full_name, email, age, division) VALUES ($1, $2, $3, $4) RETURNING *`

	err = db.QueryRow(sqlStatement, "Airell", "wWQp0@example.com", 23, "IT").Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee data: %+v\n", employee)
}

func GetEmployees() {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}
		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}
		results = append(results, employee)
	}
	fmt.Println("Employee data: ", results)
}

func UpdateEmployee(){
	sqlStatement := `UPDATE employees SET email = $2 WHERE id = $1`
	res, err := db.Exec(sqlStatement, 1, "wWQp0@example.com")

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Update employee count: ", count)
}

func DeleteEmployee(){
	sqlStatement := `DELETE FROM employees WHERE id = $1`
	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete employee count: ", count)
}