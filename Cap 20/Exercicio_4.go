/*
Exercicio 4:
- Utilize mutex para consertar a condição de corrida do exercício anterior
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

var mu sync.Mutex

func main() {

	contador := 0
	total := 10

	wg.Add(total)

	for i := 0; i < total; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println(contador)
}
