/*
Exercicio 5:
- Utilizando este código: https://play.golang.org/p/YHOMV9NYKK
- ...demonstre o comma ok idiom.
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 10
	}()

	v, ok := <-c
	fmt.Println(v, ok)
	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
