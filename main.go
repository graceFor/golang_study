package main

import (
	"github.com/graceFor/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	dictionary.Add(word, "First")
	dictionary.Update(word, "Second")

}
