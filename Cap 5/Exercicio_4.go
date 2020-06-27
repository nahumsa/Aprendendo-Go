/*
Exercicio 4:
- Crie um programa que:
    - Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    - Demonstre esta outra variável em decimal, binário e hexadecimal
*/

package main

import "fmt"

var x int = 200

func main() {
	fmt.Println("Decimal \t Binario \t Hexadecimal")
	fmt.Printf("%d \t\t %b \t\t %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("Deslocado para Esquerda\n")
	fmt.Printf("%d \t\t %b \t\t %#x\n", y, y, y)
	z := x >> 1
	fmt.Printf("Deslocado para Direita\n")
	fmt.Printf("%d \t\t %b \t\t %#x\n", z, z, z)
}
