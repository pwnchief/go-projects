package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	ans := 1
	for _, ch := range input {
		str := string(ch)
		if strings.ToUpper(str) == str {
			ans++
		}
	}
	fmt.Println(ans)
}
