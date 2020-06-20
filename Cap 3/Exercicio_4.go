/*
Exercicio 4:
- Crie um tipo. O tipo subjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
- Na função main:
    1. Demonstre o valor da variável "x"
    2. Demonstre o tipo da variável "x"
    3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x"
- Para os aventureiros: https://golang.org/ref/spec#Types
*/

package main

import (
	"fmt"
)

type wululu int

var x wululu

func main() {
	fmt.Printf("x: \t valor: %v \t tipo: %T", x, x)
	x = 42
	fmt.Printf("\nx: \t valor: %v \t tipo: %T\n", x, x)
}
