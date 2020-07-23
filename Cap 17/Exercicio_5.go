/*
Exercicio 5:
- Partindo do código abaixo, ordene os []user por idade e sobrenome.
    - https://play.golang.org/p/BVRZTdlUZ_
- Os valores no campo Sayings devem ser ordenados tambem,
  e demonstrados de maneira esteticamente harmoniosa.

*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func printH(us []user) {
	for i, v := range us {
		fmt.Printf("Personagem %v \t nome: %v %v \t idade: %v\n", i, v.First, v.Last, v.Age)
		fmt.Printf("Sayings: ")
		for index, s := range v.Sayings {
			if index == 0 {
				fmt.Printf("%v\n", s)
			} else {
				fmt.Printf("\t %v\n", s)
			}

		}

	}
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("Sem Ordenação:")
	printH(users)
	fmt.Println("")

	sort.Sort(ByAge(users))
	fmt.Println("Por Idade:")
	printH(users)
	fmt.Println("")

	sort.Sort(ByLast(users))
	fmt.Println("Por Sobrenome:")
	printH(users)
	fmt.Println("")

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	fmt.Println("Por sayings:")
	printH(users)
	fmt.Println("")

}
