/*
Exercicio 5:
- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
*/

package main

import (
	"fmt"
)

func main() {
	numDivisao := 4
	for x := 10; x <= 100; x++ {
		fmt.Printf("%v ", x%numDivisao)
	}

}
