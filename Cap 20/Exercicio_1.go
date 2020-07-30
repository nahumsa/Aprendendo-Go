/*
Exercicio 1:
- Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.

*/

package main

import (
	"fmt"
	"sync"
)

func prints(s string) {
	fmt.Println(s)
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go prints("Print 1")
	go prints("Print 2")

	wg.Wait()
}
