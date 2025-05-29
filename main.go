package main

import (
	"fmt"
	"github.com/fatih/color"
	"project-go/greeting"
)

func main() {
	fmt.Println(greeting.Hello())
	color.Cyan(greeting.Bye())
}
