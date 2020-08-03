/*
Exercicio 3:
- Utilizando este código: https://play.golang.org/p/sfyu4Is3FG
- ...use um for range loop para demonstrar os valores do canal.
*/

package main

import (
	"fmt"
)

func main() {

	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(r <-chan int) {
	for v := range r {
		fmt.Println(v)
	}

}
