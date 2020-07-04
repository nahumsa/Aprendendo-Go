/*
- Utilizando o exercício anterior,
	remova uma entrada do map e demonstre o map inteiro utilizando range.

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
	delete(data, "Baballo_Alo")
	fmt.Printf("Depois de deletar\n\n")
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
