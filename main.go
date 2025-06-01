package main

import (
	"fmt"
	"github.com/fatih/color"
	"project-go/greeting"
	"strings"
	"unicode"
)
func LatinLetters(s string) string {
	sb := strings.Builder{}

	for _, r := range s {
		if unicode.Is(unicode.Latin, r) {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
func main() {
	fmt.Println(greeting.Hello(), greeting.Bye())
	color.Cyan(greeting.Bye())
	fmt.Println(LatinLetters(" abc1"))
}
