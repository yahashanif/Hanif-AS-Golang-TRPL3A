package common

type Customer struct {
	CustomerID   string `json:"CustomerID"`
	CompanyName  string `json:"CompanyName"`
	ContactName  string `json:"ContactName"`
	ContactTitle string `json:"ContactTitle"`
	Address      string `json:"Address"`
	City         string `json:"City"`
	Country      string `json:"Country"`
	Phone        string `json:"Phone"`
	PostalCode   string `json:"PostalCode"`
}

type Employee struct {
	EmployeeID      string `json:"EmployeeID"`
	LastName        string `json:"LastName"`
	FirstName       string `json:"FirstName"`
	Title           string `json:"Title"`
	TitleOfCourtesy string `json:"TitleOfCourtesy"`
	BirthDate       string `json:"BirthDate"`
	HireDate        string `json:"HireDate"`
	Address         string `json:"Address"`
	City            string `json:"City"`
	Region          string `json:"Region"`
	PostalCode      string `json:"PostalCode"`
	Country         string `json:"Country"`
	HomePhone       string `json:"HomePhone"`
	Extension       string `json:"Extension"`
	Photo           string `json:"Photo"`
	Notes           string `json:"Notes"`
	ReportsTo       string `json:"ReportsTo"`
	ProvinceName    string `json:"ProvinceName"`
}
