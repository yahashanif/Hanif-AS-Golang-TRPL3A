package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Order struct (Model) ...
type Orders struct {
	OrderID      string        `json:"orderID"`
	CustomerID   string        `json:"customerID"`
	EmployeeID   string        `json:"employeeID"`
	OrderDate    string        `json:"orderDate"`
	OrdersDet []OrdersDetail   `json:"OrdersDet"`
	
}

type OrdersDetail struct {
	OrderID      string  `json:"orderID"`
	ProductID  	 string  `json:"ProductID"`
	ProductName  string  `json:"ProductName"`
	UnitPrice    float64 `json:"UnitPrice"`
	Quantity     int     `json:"Quantity"`
}

func main() {
	
	url := "http://localhost:8080/orders"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	
	if readErr != nil {
		log.Fatal(readErr)
	}

	orders := Orders{}
	jsonErr := json.Unmarshal(body, &orders)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(orders.CustomerID)
	fmt.Println(orders.EmployeeID)
	fmt.Println(orders.OrderID)

	for _, product := range orders.OrdersDet {
		fmt.Println("Order ID",product.OrderID) 
		fmt.Println("Product ID",product.ProductID)
		fmt.Println("Product Name",product.ProductName)
		fmt.Println("Quantity",product.Quantity)
		fmt.Println("Unit Price",product.UnitPrice)
	}
	
	
	
}