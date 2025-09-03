package main

import "fmt"

func main() {
		var age int
		fmt.Println("Qual sua idade?")
		fmt.Scanln(&age)
		if age >= 18{
			fmt.Println("Vc é maior, meu chapa!")
		}	else {
			fmt.Println("Vc é menor meu amigo.")
		}
}