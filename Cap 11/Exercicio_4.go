/*
Exercicio 4:
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/

package main

import "fmt"

func main() {

	x := struct {
		dicionario map[string]int
		tipos      []string
	}{
		dicionario: map[string]int{
			"Teste1": 90,
			"Teste2": 95,
			"Teste3": 99,
		},

		tipos: []string{
			"Prova dificil",
			"Prova média",
			"Prova fácil",
		},
	}

	fmt.Println("Dicionario:", x.dicionario)
	fmt.Println("Tipos:", x.tipos)

}
