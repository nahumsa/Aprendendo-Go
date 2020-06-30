/*
Exercicio 2:
- Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
- Por exemplo:
    65
        U+0041 'A'
        U+0041 'A'
        U+0041 'A'
    66
        U+0042 'B'
        U+0042 'B'
        U+0042 'B'
*/

package main

import (
	"fmt"
)

func main() {

	minUnicode := 65
	maxUnicode := 90
	repeticao := 3
	for x := minUnicode; x <= maxUnicode; x++ {
		fmt.Printf("%v \n", x)
		for y := 0; y < repeticao; y++ {
			fmt.Printf("\t%U %v \n", x, string(x))
		}
	}

}
