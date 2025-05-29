package main

import (
  "fmt"
  "project-go/greeting"
  "github.com/fatih/color"
)
func main() {
  fmt.Println(greeting.Hello())
  color.Cyan(greeting.Bye())
}
