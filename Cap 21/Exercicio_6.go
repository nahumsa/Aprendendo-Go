/*
Exercicio 6:
- Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.
*/

package main

import "fmt"

func main() {
	quit := make(chan bool)
	c := create(100, quit)

	show(c, quit)

	fmt.Println("Program has ended.")
}

func show(c <-chan int, q chan bool) {
	for {
		select {
		case v := <-c:
			fmt.Println("Valor:", v)

		case <-q:
			fmt.Println("Quit!")
			return
		}

	}
}

func create(n int, q chan bool) <-chan int {
	c := make(chan int)

	go func() {
		for i := 1; i < n+1; i++ {
			c <- i
		}
		q <- true
		close(c)
	}()

	return c
}
