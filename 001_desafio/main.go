package main

import (
	"fmt"
	"time"
)

func main() {

	currentYear := time.Now().Year()
	var yearOfBirth int

	for {
		var age int

		fmt.Println("Qual a sua idade?")

		_, err := fmt.Scan(&age)

		if err != nil {
			fmt.Println("Entrada inválida, por favor repita")
			continue
		}

		yearOfBirth = currentYear - age
		break
	}

	fmt.Printf("Seu ano de nascimento é %d\n", yearOfBirth)
}
