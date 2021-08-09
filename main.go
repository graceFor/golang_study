package main

import "fmt"

func main() {
	a := 2
	b := &a // a의 주소를 복사
	*b = 20
	fmt.Println(*a)
}
