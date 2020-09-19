/*
Exercicio 2:
- Coloque seu código no GitHub.
- Faça sua documentação aparecer em godoc.org, e tire um screenshot.
- Delete seu código do GitHub.
- Faça um refresh em godoc.org e veja se seu código sumiu.
- Me mostre no twitter (@ellenkorbes) seu screenshot!
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
