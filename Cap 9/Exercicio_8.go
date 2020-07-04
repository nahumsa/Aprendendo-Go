/*
Exercicio 8:
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.

*/

package main

import "fmt"

func main() {
	data := map[string][]string{
		"Alibaba_Babo": []string{"Algo", "Blabla"},
		"Baballo_Alo":  []string{"Nada"},
	}

	for i, val := range data {
		fmt.Printf("Nome: %v\n", i)
		fmt.Printf("Hobbies: ")
		for index, hobby := range val {
			fmt.Printf("%v - %v\n", index, hobby)
			fmt.Printf("\t ")
		}
		fmt.Printf("\n\n")

	}

}
