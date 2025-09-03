package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Digite um número: ")
	fmt.Scanln(&num1)

	fmt.Println("Digite outro número: ")
	fmt.Scanln(&num2)

	sum := num1+num2

	fmt.Printf("A soma de %d e %d é igual a %d \n", num1, num2, sum)
}