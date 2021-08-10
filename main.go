package main

import (
	"fmt"

	"github.com/graceFor/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("hyun")
	//fmt.Println(account)  // &{hyun 0}
	//fmt.Println(*account) // {hyun 0}
	account.Despite(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
