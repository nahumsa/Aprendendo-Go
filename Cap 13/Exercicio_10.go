/*
Exercicio 10:

- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função,
  onde esta outra função faz uso de uma variável alem de seu scope.

*/

package main

import "fmt"

func somaPorChamada() func() int {
	x := 0
	defer fmt.Printf("Funcao Chamada\n")
	return func() int {
		x++
		return x
	}
}

func main() {
	a := somaPorChamada()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println("")

	b := somaPorChamada()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
