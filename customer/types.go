package customer

const dataFile = "customers.json"

type customerData struct {
	Customers  []Customer `json:"customers"`
	CustomerId int        `json:"customerId"`
}

var customers []Customer
var customerId int = 0

type Customer struct {
	Cid    int
	Name   string
	Age    int
	Gender string
	Email  string
}
