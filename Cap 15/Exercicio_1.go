/*
Exercicio 1:
- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.

*/

package main

import "fmt"

func main() {
	x := "Memória"

	fmt.Println("Valor de x na memória: ", &x)
}
