/*
Exercicio 2:
- Crie um struct "pessoa"
- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. ]
  Essa função deve mudar um valor armazenado no endereço *pessoa.
	- Dica: a maneira "correta" para fazer dereference de um valor em um struct seria
	  (*valor).campo
    - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
	- "As an exception, if the type of x is a named pointer type and (*x).f
	  is a valid selector expression denoting a field (but not a method),
    - → x.f is shorthand for (*x).f." ←
    - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
- Crie um valor do tipo pessoa;

- Use a função mudeMe e demonstre o resultado.

*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func maisVelho(p *pessoa) {
	(*p).idade++
}

func casou(p *pessoa, novoNome string) {
	(*p).sobrenome = novoNome
}

func main() {
	x := pessoa{"Agnobaldo", "Agnácio", 30}
	fmt.Println("Idade: ", x.idade)
	fmt.Printf("Parabéns, %v !\n", x.nome)
	maisVelho(&x)
	fmt.Println("Idade: ", x.idade)

	y := pessoa{"Bernado", "Burlo", 30}
	fmt.Printf("%v %v casou com %v %v\n", y.nome, y.sobrenome, x.nome, x.sobrenome)
	casou(&y, x.sobrenome)
	fmt.Printf("Agora o nome dele é: %v %v!\n", y.nome, y.sobrenome)

}
