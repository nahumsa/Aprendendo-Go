/*
Exercicio 6:
- Crie um programa que demonstra seu OS e ARCH.
- Rode-o com os seguintes comandos:
    - go run
    - go build
	- go install

*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Sistema Operacional:", runtime.GOOS)
	fmt.Println("Arquitetura:", runtime.GOARCH)
}
