package main

import "fmt"

// func main() {
// 		for i:= 0; i <= 5; i++ {
// 			fmt.Printf("Essa eh a iteracao numer %d\n", i)
// 		}
// }


func main() {
	  var age int
		for {
			fmt.Printf("Por favor diga a sua idade: \n")
			_,err := fmt.Scanln(&age)
			if err!= nil {
				fmt.Println("Entrada inválida, por favor repita")
				continue
			}

			if age < 18 {
				fmt.Println("Entrada inválida, por favor repita")
				continue
			}
			break
		}

		fmt.Println("Seja bem vindo")
}