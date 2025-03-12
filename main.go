package main

import "fmt"

func main() {

	var customer1 = map[string]string{"name": "yuan", "age": "22"}
	var customer2 = map[string]string{"name": "fang", "age": "23"}
	var customer3 = map[string]string{"name": "quan", "age": "24"}

	var customers = []map[string]string{customer1, customer2, customer3}

	var newCustomer = map[string]string{"name": "xiaoqiang", "age": "35"}

	//增
	customers = append(customers, newCustomer)
	
	//删
	customers = append(customers[:2], customers[3:]...)

	//改
	customers[2]["name"] = "xiaoqiang"

	//查
	fmt.Println(customers)
	fmt.Printf("%s\n", customers[2]["name"])

}
