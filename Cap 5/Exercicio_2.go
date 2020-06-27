/*
Exercicio 2:
- Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
	- ==
	- !=
	- <=
	- <
	- >=
	- >

- Demonstre estes valores.
*/

package main

import (
	"fmt"
)

var x = 10
var y = 50

func main() {
	igual := (x == y)
	nigual := (x != y)
	menorig := (x <= y)
	menor := (x < y)
	maiorig := (x >= y)
	maior := (x > y)
	fmt.Printf("x= %v, y= %v\n", x, y)
	fmt.Printf("Igual: %v\n", igual)
	fmt.Printf("Diferente: %v\n", nigual)
	fmt.Printf("Menor ou igual: %v\n", menorig)
	fmt.Printf("Menor: %v\n", menor)
	fmt.Printf("Maior ou igual: %v\n", maiorig)
	fmt.Printf("Maior: %v\n", maior)
}
