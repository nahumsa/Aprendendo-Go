/*
Exercicio 3:
- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
    - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
    - Do quinto ao último item do slice (incluindo o último item!)
    - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
    - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
    - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
*/

package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("Indice 1 ao 3: %v\n", array[:3])
	fmt.Printf("Indice 5 ao ultimo: %v\n", array[4:])
	fmt.Printf("Indice 2 ao 7: %v\n", array[1:7])
	fmt.Printf("Indice 3 ao penultimo: %v\n", array[2:9])
	fmt.Printf("Indice 3 ao penultimo utilizando len: %v\n", array[2:len(array)-1])
}
