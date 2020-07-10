/*
Exercicio 1:
	- Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
*/

package main

import "fmt"

func soma(x ...int) int {
	soma := 0
	for _, val := range x {
		soma += val
	}
	return soma
}

func saborFavorito(sabor string) (string, int) {
	return fmt.Sprintf("Sabor favorito: %v", sabor), 1
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Soma: %v \n", soma(numeros...))
	sabor1, numeroSabor := saborFavorito("Canela")

	fmt.Println(sabor1, "\nQuantos sabores?", numeroSabor)
}
