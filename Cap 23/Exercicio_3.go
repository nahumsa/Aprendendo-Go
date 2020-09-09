package main

import "fmt"

/*
Exercicio 3:
	- Crie um struct "erroEspecial" que implemente a interface builtin.error.
	- Crie uma função que tenha um valor do tipo error como parâmetro.
	- Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.
*/

type erroEspecial struct {
	nome string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("%v", e.nome)
}

func deuRuim(e error) {
	fmt.Println(e.Error())
}

func oops() error {
	return erroEspecial{"Opa! Deu ruim!"}
}

func main() {
	err := oops()

	if err != nil {
		deuRuim(err)
	}
}
