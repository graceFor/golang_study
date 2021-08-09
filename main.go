package main

import (
	"fmt"
	"strings"
)

func lenAndupper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// func multiply(a int, b int) int {
// 	return a * b
// }
func multiply(a, b int) int {
	return a * b
}

// ...은 내가 원하는 만큼 파라미터를 전달하는 방법
func repeatMe(words ...string) {
	fmt.Println((words))
}

func main() {
	repeatMe("hyun", "kyung", "kim", "g")

	// 상수  = 값 변경 X
	const name string = "hyun"
	fmt.Println(name)

	// 변수  = 값 변경 O
	//var name1 string = "kyung"
	name1 := "kyung"
	name1 = "kyungs"
	fmt.Println((name1))

	fmt.Println(multiply(3, 4))
	totalLength, _ := lenAndupper("hyunkyung")
	fmt.Println(totalLength)

}
