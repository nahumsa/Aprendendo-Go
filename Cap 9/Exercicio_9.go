/*
Exercicio 9:
- Utilizando o exercício anterior, adicione uma entrada ao map e
demonstre o map inteiro utilizando range.

*/

package main

import "fmt"

func main() {
	data := map[string][]string{
		"Alibaba_Babo": []string{"Algo", "Blabla"},
		"Baballo_Alo":  []string{"Nada"},
	}

	data["Carlos_Carlas"] = []string{"Alguma Coisa"}

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
