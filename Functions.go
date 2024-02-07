package main

import (
	"fmt"
)

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)
	fmt.Println("**********************")
	Area, Parameter := rectProps(10.8, 5.6)
	fmt.Println("Area", Area, "Parameter", Parameter)
}
