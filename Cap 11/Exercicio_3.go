/*
Exercicio 3:
- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.
*/

package main

import "fmt"

type veículo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veículo
	traçãoNasQuatro bool
}

type sedan struct {
	veículo
	modeloLuxo bool
}

func main() {

	Caminhonete := caminhonete{
		veículo:         veículo{2, "Preto"},
		traçãoNasQuatro: true,
	}

	Sedan := sedan{
		veículo:    veículo{4, "Cinza"},
		modeloLuxo: true,
	}
	fmt.Println("Sedan:", Sedan)
	fmt.Println("É de luxo?", Sedan.modeloLuxo)
	fmt.Println("Caminhonete:", Caminhonete)
	fmt.Println("Tem tração nas quatro rodas?", Caminhonete.traçãoNasQuatro)

}
