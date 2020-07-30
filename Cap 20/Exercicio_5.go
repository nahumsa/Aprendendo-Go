/*
Exercicio 5:
- Utilize atomic para consertar a condição de corrida do exercício #3.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var contador int64
	contador = 0
	total := 10

	wg.Add(total)

	for i := 0; i < total; i++ {
		go func() {

			atomic.AddInt64(&contador, 1)
			runtime.Gosched()

			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println(contador)
}
