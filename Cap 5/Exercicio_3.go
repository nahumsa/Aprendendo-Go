/*
Exercicio 3:
- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.
*/

package main

import "fmt"

const x int = 10 // Constante tipada
const y = 10.    // Constante não-tipada

func main() {
	fmt.Printf("x:\n  valor: %v tipo:%T\n", x, x)
	fmt.Printf("y:\n  valor: %v tipo:%T\n", y, y)
}
