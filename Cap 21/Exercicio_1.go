/*
Exercicio 1:
- Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
    - Usando uma função anônima auto-executável
    - Usando buffer
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()
	// v := <-c

	fmt.Println("Sem buffer:", <-c)

	d := make(chan int, 1)

	go func() {
		d <- 42
	}()

	fmt.Println("Com buffer:", <-d)

}
