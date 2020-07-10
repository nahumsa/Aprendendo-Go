/*
Exercicio 9:

- Callback: passe uma função como argumento a outra função.

*/

package main

import "fmt"

func quadrado(numero int) int {
	return numero * numero
}

func somaFuncao(f func(x int) int, y ...int) int {
	total := 0
	for _, val := range y {
		total += f(val)
	}
	return total
}

func main() {
	exemplo := somaFuncao(quadrado, []int{1, 2, 3, 4, 5, 6}...)
	fmt.Println(exemplo)
}
