package main

import (
	"fmt"
	"strings"
)

// func lenAndupper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func lenAndupper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {

	totalLength, upper := lenAndupper("hyunkyung")
	fmt.Println(totalLength, upper)

}
