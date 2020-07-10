/*
Exercicio 7:
- Atribua uma função a uma variável.
- Chame essa função.
*/

package main

import (
	"fmt"
)

func main() {

	quadrado := func(numero int) int {
		return numero * numero
	}

	x := 2

	fmt.Println(quadrado(x))

}
