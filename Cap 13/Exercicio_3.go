/*
Exercicio 3:
- Utilize a declaração defer de maneira que demonstre que
sua execução só ocorre ao final do contexto ao qual ela pertence.
*/

package main

import "fmt"

func main() {

	fmt.Println("Sem Defer: \n")
	fmt.Println("Ultimo")
	fmt.Println("Primeiro")

	fmt.Println("\nCom Defer: \n")
	defer fmt.Println("Ultimo")
	fmt.Println("Primeiro")
}
