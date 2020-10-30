package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	cm "Hanif-AS-Golang-TRPL3A/HtmlPage/common"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Index(w http.ResponseWriter, r *http.Request) {

	var customers []cm.Customer

	sql := `SELECT
				CustomerID,
				IFNULL(CompanyName,''),
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Country,'') Country,
				IFNULL(Phone,'') Phone ,
				IFNULL(PostalCode,'') PostalCode
			FROM customers ORDER BY CustomerID`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var customer cm.Customer
		err := result.Scan(&customer.CustomerID, &customer.CompanyName, &customer.ContactName,
			&customer.ContactTitle, &customer.Address, &customer.City, &customer.Country,
			&customer.Phone, &customer.PostalCode)

		if err != nil {
			panic(err.Error())
		}
		customers = append(customers, customer)
	}

	t, err := template.ParseFiles("index.html")
	t.Execute(w, customers)

	if err != nil {
		panic(err.Error())
	}

}

func IndexEmployee(w http.ResponseWriter, r *http.Request) {

	var employee []cm.Employee

	sql := `SELECT
				EmployeeID,
				IFNULL(LastName,''),
				IFNULL(FirstName,'') FirstName,
				IFNULL(Title,'') Title,
				IFNULL(TitleOfCourtesy,'') TitleOfCourtesy,
				IFNULL(BirthDate,'') BirthDate,
				IFNULL(HireDate,'') HireDate,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Region,'') Region,
				IFNULL(PostalCode,'') PostalCode,
				IFNULL(Country,'') Country,
				IFNULL(HomePhone,'') HomePhone,
				IFNULL(Extension,'') Extension,
				IFNULL(Photo,'') Photo,
				IFNULL(Notes,'') Notes,
				IFNULL(ReportsTo,'') ReportsTo,
				IFNULL(ProvinceName,'') ProvinceName

				
			FROM employees order by EmployeeID`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var employees cm.Employee
		err := result.Scan(&employees.EmployeeID, &employees.LastName,
			&employees.FirstName, &employees.Title, &employees.TitleOfCourtesy,
			&employees.BirthDate, &employees.HireDate, &employees.Address,
			&employees.City, &employees.Region, &employees.PostalCode,
			&employees.Country, &employees.HomePhone, &employees.Extension,
			&employees.Photo, &employees.Notes, &employees.ReportsTo,
			&employees.ProvinceName)

		if err != nil {
			panic(err.Error())
		}
		employee = append(employee, employees)
	}

	t, err := template.ParseFiles("indexemployee.html")
	t.Execute(w, employee)

	if err != nil {
		panic(err.Error())
	}

}

func main() {
	//<user>:<passwprd>@tcp<IP address>/<Password>
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/northwind")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	log.Println("Server started on: http://localhost:8081")
	http.HandleFunc("/", Index)
	http.HandleFunc("/employee", IndexEmployee)
	http.ListenAndServe(":8081", nil)

}
