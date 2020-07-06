/*
Exercicio 2:
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map,
  utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range,
  dentro do range anterior.
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

	TodosClientes := map[string]pessoa{
		Cliente1.Sobrenome: Cliente1,
		Cliente2.Sobrenome: Cliente2,
	}

	for sobrenome, cliente := range TodosClientes {
		fmt.Printf("Sobrenome: %v\n", sobrenome)
		fmt.Printf("\tSabores de Sorvete:")
		for _, sabores := range cliente.SaboresSorvete {
			fmt.Printf("%v\t", sabores)
		}
		fmt.Printf("\n")
	}

}
