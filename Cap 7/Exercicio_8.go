/*
Exercicio 8:
- Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/

package main

import "fmt"

func main() {
	numero := 7

	switch {
	case numero%2 == 0 && numero%3 == 0:
		fmt.Println("Divisivel por 2 e por 3")
	case numero%2 == 0:
		fmt.Println("Divisivel por 2")
	case numero%3 == 0:
		fmt.Println("Divisivel por 3")
	default:
		fmt.Println("Não é divisivel por 2 nem por 3")
	}
}
