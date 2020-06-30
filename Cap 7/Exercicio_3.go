/*
Exercicio 3:
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import (
	"fmt"
)

func main() {

	anoNascimento := 1994
	anoAtual := 2020
	for anoNascimento <= anoAtual {
		fmt.Printf("%v \n", anoNascimento)
		anoNascimento++
	}

}
