package main

import (
	"fmt"

	"github.com/graceFor/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	word, _ := dictionary.Search(baseword)
	fmt.Println(word)
	dictionary.Delete(baseword)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
