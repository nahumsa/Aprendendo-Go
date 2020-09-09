/*
Exercicio 5:

- Nos capítulos seguintes, uma das coisas que veremos é testes.
- Para testar sua habilidade de se virar por conta própria... desafio:
    - Utilizando as seguintes fontes: https://godoc.org/testing & http://www.golang-book.com/books/intr...
    - Tente descobrir por conta própria como funcionam testes em Go.
    - Pode usar tradutor automático, pode rodar código na sua máquina, pode procurar no Google. Vale tudo.
    - O exercício é: crie um teste simples de uma função ou método ou pedaço qualquer de código.
*/

package Exercicio5

import (
	"fmt"
	"testing"
)

// TestSum do a simple test on the function sum
func TestSum(t *testing.T) {
	want := 45
	got := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	if want != got {
		fmt.Errorf("Expected: %v, Got: %v", want, got)
	}
}

func sum(list []int) int {
	varSum := 0
	for _, val := range list {
		varSum += val
	}
	return varSum
}
