package main

import (
	f "fmt"
)

var name = "FEM"

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func recalculateTax(price float32) []float32 {
	return []float32{price * 0.09, price * 0.02}
}

func xrecalculateTax(price float32) (stateTax float32, cityTax float32) {
	stateTax = price * 0.09
	cityTax = price * 0.02
	return stateTax, cityTax
}

func birthday(age *int) {
	if *age >= 30 {
		panic("Too old to be true")
	}
	f.Printf("The pointer is %v and the value is %v\n", age, *age)
	*age++
}

func main() {
	//f.Println(xrecalculateTax(2.54))
	//f.Println(recalculateTax(2.3))
	//stateTax, _ := calculateTax(100)
	//f.Println(stateTax)

	defer f.Println("Bye!!")
	age := 25
	birthday(&age)
	birthday(&age)
	birthday(&age)
	birthday(&age)
	birthday(&age)
	f.Println(age)
	//PrintData()
}
