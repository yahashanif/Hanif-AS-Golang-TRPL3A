package function

import (
	"database/sql"
	"html/template"
	"net/http"
)

var db *sql.DB
var err error

func RouteIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("index.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func RouteSubmitPost(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/northwind")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	if r.Method == "POST" {

		var FirstName = r.FormValue("firstname")
		var LastName = r.Form.Get("lastname")
		var Title = r.FormValue("Title")
		var TitleOfCourtesy = r.FormValue("TitleOfCourtesy")
		var BirthDate = r.FormValue("BirthDate")
		var HireDate = r.FormValue("HireDate")
		var Address = r.FormValue("Address")
		var City = r.FormValue("City")
		var Region = r.FormValue("Region")
		var PostalCode = r.FormValue("PostalCode")
		var Country = r.FormValue("Country")
		var HomePhone = r.FormValue("HomePhone")
		var Extension = r.FormValue("Extension")
		Photo := r.FormValue("Photo")
		Notes := r.FormValue("Notes")
		ReportsTo := r.FormValue("ReportsTo")
		ProvinceName := r.FormValue("ProvinceName")

		stmt, err := db.Prepare("INSERT INTO employees VALUES (NULL,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

		_, err = stmt.Exec(LastName, FirstName, Title,
			TitleOfCourtesy, BirthDate, HireDate, Address,
			City, Region, PostalCode, Country,
			HomePhone, Extension, Photo, Notes, ReportsTo,
			ProvinceName)

		var tmpl = template.Must(template.New("result").ParseFiles("index.html"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			var data = map[string]string{"firstname": FirstName,
				"lastname":        LastName,
				"title":           Title,
				"TitleOfCourtesy": TitleOfCourtesy,
				"BirthDate":       BirthDate,
				"HireDate":        HireDate,
				"Address":         Address,
				"City":            City,
				"Region":          Region,
				"PostalCode":      PostalCode,
				"Country":         Country,
				"HomePhone":       HomePhone,
				"Extension":       Extension,
				"Photo":           Photo,
				"Notes":           Notes,
				"ReportsTo":       ReportsTo,
				"ProvinceName":    ProvinceName}

			if err := tmpl.Execute(w, data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		if err := r.ParseForm(); err != nil {

		}

		//tugas insertkan ke database ke table user

		return

		http.Error(w, "", http.StatusBadRequest)

	}

}
