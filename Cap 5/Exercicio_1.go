/*
Exercicio 1:
- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
*/

package main

import (
	"fmt"
)

func main() {
	numero := 100
	fmt.Println("Decimal \t Binario \t Hexadecimal")
	fmt.Printf("%d \t\t %b \t\t %#x\n", numero, numero, numero)
}
