/*
Exercicio 8:
- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.

*/

package main

import (
	"fmt"
)

func retornaQuadrado() func() {

	fmt.Println("Primeira Função\n")
	return func() {
		fmt.Println("Segunda Função")
	}
}

func main() {

	quad := retornaQuadrado()

	fmt.Println("Chamando a segunda função\n")
	quad()

}
