/*
Exercicio 4:
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import (
	"fmt"
)

func main() {

	anoNascimento := 1994
	anoAtual := 2020
	for {
		fmt.Printf("%v \n", anoNascimento)
		anoNascimento++
		if !(anoNascimento <= anoAtual) {
			break
		}
	}

}
