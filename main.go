package main

import (
	"fmt"

	"github.com/graceFor/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeating"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	searchDefinition, _ := dictionary.Search(word)
	fmt.Println("found:", word, ",definition:", searchDefinition)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

}
