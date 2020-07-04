/*
Exercicio 7:
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/

package main

import "fmt"

func main() {
	dados := [][]string{
		[]string{"Alibaba", "Babo", "Algo"},
		[]string{"Baballo", "Alo", "Nada"},
		[]string{"Carlos", "Carlas", "Alguma Coisa"},
	}
	fmt.Println("Nome\tSobrenome Hobby")
	for _, val := range dados {
		for _, dat := range val {
			fmt.Printf("%v\t", dat)
		}
		fmt.Printf("\n")

	}
}
