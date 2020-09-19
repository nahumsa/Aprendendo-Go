/*
Exercicio 1:
- Crie um package "cachorro".
	- Este package deverá exportar uma função Idade, que toma como parâmetro um número de anos e
	retorna a idade equivalente em anos caninos. (1 ano humano → 7 anos caninos)
    - Documente seu código com comentários, e utilize a função Idade na sua função main.
- Rode seu programa para verificar se ele funciona.
- Rode um local server com godoc e leia sua documentação.

*/

// Package cachorro exporta uma função representa a idade humana em idade canina.
package cachorro

import "fmt"

// Idade converte de anos humanos para anos caninos
func Idade(anosHumanos int) int {
	return anosHumanos * 7
}

// main testa a função
func main() {
	fmt.Println(Idade(10))
}
