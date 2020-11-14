package main

import "fmt"

func main() {
	fmt.Println("Please enter First number :")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("Please enter Second number :")
	var num2 int
	fmt.Scanln(&num2)
	var sum int
	sum = num1+num2
	fmt.Println("The sum is ",sum)
} 