/*
Exercicio 6:
- Crie um programa que demonstre o funcionamento da declaração if.
*/

package main

import (
	"fmt"
)

func main() {
	numDivisao := 2
	fmt.Println("Numeros divisíveis por ", numDivisao)
	for x := 10; x <= 100; x++ {
		if x%numDivisao == 0 {
			fmt.Printf("%v ", x)
		}

	}

}
