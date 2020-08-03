/*
Exercicio 7:
- Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
- Tire estes números do canal e demonstre-os.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	canal := make(chan int)
	var wg sync.WaitGroup
	n := 10

	for i := 0; i < n; i++ {
		wg.Add(1)
		go envia(canal, n)
		wg.Done()
	}

	wg.Wait()

	for j := 0; j < n*n; j++ {
		fmt.Println(<-canal)
	}
}

func envia(canal chan int, n int) {
	for j := 0; j < n; j++ {
		canal <- j
	}
}
