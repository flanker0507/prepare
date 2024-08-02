package main

import (
	"fmt"
	"strings"
)

func main() {
	var index1 = strings.Index("ethan hunt", "e")
	fmt.Println(index1)

	text := "yuda ganteng"
	find := "z"
	replace := "d"

	newText1 := strings.Replace(text, find, replace, 4)
	fmt.Println(newText1)
}
