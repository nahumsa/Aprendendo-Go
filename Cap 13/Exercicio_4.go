/*
Exercicio 4:
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.

*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) printNome() {

	fmt.Printf("Nome Completo: %v %v \nIdade: %v\n", p.nome, p.sobrenome, p.idade)
}

func main() {
	pessoa1 := pessoa{"Arnaldo", "Agnobaldo", 70}
	pessoa1.printNome()
	pessoa2 := pessoa{"Arlingo", "Baldo", 20}
	pessoa2.printNome()
}
