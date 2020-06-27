/*
Exercicio 6:
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
*/

package main

import "fmt"

const (
	x = iota + 2020
	y
	z
	a
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
}
