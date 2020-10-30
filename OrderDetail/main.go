package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Order struct (Model) ...
type Orders struct {
	OrderID      string        `json:"orderID"`
	CustomerID   string        `json:"customerID"`
	EmployeeID   string        `json:"employeeID"`
	OrderDate    string        `json:"orderDate"`
	OrdersDet []OrdersDetail   `json:"ordersDetail"`
	
}

type OrdersDetail struct {
	OrderID      string  `json:"orderID"`
	ProductID  	 string  `json:"ProductID"`
	ProductName  string  `json:"ProductName"`
	UnitPrice    float64 `json:"UnitPrice"`
	Quantity     int     `json:"Quantity"`
}

// Get all orders

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


	var order Orders	
	var orderdet OrdersDetail	

	sql := `SELECT
				OrderID,
				IFNULL(CustomerID,'') CustomerID,
				IFNULL(EmployeeID,'') EmployeeID,
				IFNULL(OrderDate,'') OrderDate				
			FROM orders WHERE OrderID IN (11024)`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&order.OrderID,  &order.CustomerID, &order.EmployeeID, &order.OrderDate)

		if err != nil {
			panic(err.Error())
		}

		sqlDetial := `SELECT
						order_details.OrderID		
						, products.ProductID
						, products.ProductName
						, order_details.UnitPrice
						, order_details.Quantity
					FROM
						order_details
						INNER JOIN products 
							ON (order_details.ProductID = products.ProductID)
					WHERE order_details.OrderID	= ?`

		orderID := &order.OrderID
		fmt.Println(*orderID)
		resultDetail, errDet := db.Query(sqlDetial, *orderID)

		defer resultDetail.Close()

		if errDet != nil {
			panic(err.Error())
		}

		for resultDetail.Next() {
			
			err := resultDetail.Scan(&orderdet.OrderID, &orderdet.ProductID, &orderdet.ProductName, &orderdet.UnitPrice, &orderdet.Quantity)

			if err != nil {
				panic(err.Error())
			}

			order.OrdersDet = append(order.OrdersDet, orderdet)	
			
		}
				
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)

}


// Main function
func main() {

	db, err = sql.Open("mysql", "root:nadipw@tcp(127.0.0.1:3306)/northwind")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/orders", getOrders).Methods("GET")
	
	fmt.Println("Server on :8080")
	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
	
}
