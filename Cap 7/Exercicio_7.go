/*
Exercicio 7:
- Utilizando a solução anterior, adicione as opções else if e else.
*/

package main

import (
	"fmt"
)

func main() {
	numDivisao := 2
	numDivisao2 := 3
	fmt.Printf("Numeros divisíveis por %v ou %v: \n ", numDivisao, numDivisao2)
	for x := 10; x <= 100; x++ {
		if x%numDivisao == 0 {
			fmt.Printf("%v ", x)
		} else if x%numDivisao2 == 0 {
			fmt.Printf("%v ", x)
		} else {
			fmt.Println("\nNão é divisível por nenhum dos dois ", x)
		}

	}

}
