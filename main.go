package main

import (
	"fmt"

	"github.com/graceFor/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("hyun")
	fmt.Println(account)  // &{hyun 0}
	fmt.Println(*account) // {hyun 0}
}
