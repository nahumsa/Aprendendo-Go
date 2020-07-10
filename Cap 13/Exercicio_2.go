/*
Exercicio 2:
- Crie uma função que receba um parâmetro variádico do tipo int e
  retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e
  retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.

*/

package main

import "fmt"

func soma(x ...int) int {
	valor := 0

	for _, val := range x {
		valor += val
	}

	return valor
}

func somaSlice(x []int) int {
	valor := 0

	for _, val := range x {
		valor += val
	}

	return valor
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Soma: %v \n", soma(numeros...))
	fmt.Printf("SomaSlice: %v \n", somaSlice(numeros))
}
