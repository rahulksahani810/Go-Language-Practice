package main

import (
	"fmt"
)

func DeletionInMap() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)

}
func GetValueAsKey() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	employee := "jamie"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)
}

func main() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	employeeSalary["mike"] = 9000

	fmt.Println("employeeSalary map contents:", employeeSalary)
	fmt.Println("***********************")
	GetValueAsKey()
	fmt.Println("***********************")
	DeletionInMap()
}
