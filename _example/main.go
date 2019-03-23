package main

import (
	"fmt"

	"github.com/kjfsm/pokemon-go/pokemon"
)

func main() {
	happinas, err := pokemon.New(&pokemon.Config{Name: "ハピナス", Ability: "しぜんかいふく"})
	if err != nil {
		fmt.Println(err)
		return
	}
	happinas.SetLevel(50)
	fmt.Println(happinas.GetName())
	happinas.SetIV([6]int{31, 31, 31, 31, 31, 31})
	happinas.SetEV([6]int{252, 0, 0, 0, 252, 0})
	status := happinas.CalcStatus()
	fmt.Println(status)
}
