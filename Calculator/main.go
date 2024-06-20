package main

import "fmt"

func main() {
	var operations string
	var number1, number2 int
	fmt.Println("Calcualtor Go")
	fmt.Println("=============")
	fmt.Println("Which operations do you want to perform? (add, substract, multiply, divide)")
	fmt.Scanf("%s", &operations)
	fmt.Println("Enter first number")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter second number")
	fmt.Scanf("%d", &number2)
	switch operations {
	case "add":
		fmt.Println(number1 + number2)
	case "substract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "devide":
		fmt.Println(number1 / number2)
	}
}
