/*
Exercicio 10:
- Anote (à mão) o resultado das expressões:
    - fmt.Println(true && true)
    - fmt.Println(true && false)
    - fmt.Println(true || true)
    - fmt.Println(true || false)
    - fmt.Println(!true)
*/

package main

import "fmt"

func main() {
	x := true
	y := false
	fmt.Println("true && true:", x && x)
	fmt.Println("true && false:", x && y)
	fmt.Println("true || true:", x || x)
	fmt.Println("true || false:", x || y)
	fmt.Println("!true:", !x)

}
