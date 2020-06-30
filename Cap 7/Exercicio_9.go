/*
Exercicio 9:
- Crie um programa que utilize a declaração switch,
onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/

package main

import "fmt"

func main() {
	esporteFavorito := "Futebol Americano"

	switch esporteFavorito {
	case "Futebol Americano":
		fmt.Println("O Esporte favorito é:", esporteFavorito)
	case "Vôlei":
		fmt.Println("O Esporte favorito é:", esporteFavorito)
	case "Tênis":
		fmt.Println("O Esporte favorito é:", esporteFavorito)
	default:
		fmt.Println("Não tem um esporte favorito.")
	}
}
