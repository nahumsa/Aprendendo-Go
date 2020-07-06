/*
Exercicio 1:
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/

package main

import "fmt"

type pessoa struct {
	Nome           string
	Sobrenome      string
	SaboresSorvete []string
}

func main() {

	Cliente1 := pessoa{Nome: "João",
		Sobrenome:      "Estilinaldo",
		SaboresSorvete: []string{"Creme", "Flocos"}}

	Cliente2 := pessoa{Nome: "Maria",
		Sobrenome:      "Joaquina",
		SaboresSorvete: []string{"Pipoca", "Caramelo"}}

	TodosClientes := []pessoa{Cliente1, Cliente2}

	fmt.Printf("Clientes:\n")

	for i, cliente := range TodosClientes {
		fmt.Printf("%v\n\t", i)
		fmt.Printf("Nome: %v %v\n", cliente.Nome, cliente.Sobrenome)
		fmt.Printf("\tSorvetes:\n\t\t")
		for _, sabores := range cliente.SaboresSorvete {
			fmt.Printf("%v\t", sabores)
		}
		fmt.Printf("\n")

	}

}
