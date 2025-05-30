package main

import (
	"fmt"
	"github.com/fatih/color"
	"project-go/greeting"
)
func UniqueUserIDs(userIDs []int64) []int64 {
  if len(userIDs) <= 1 {
    return userIDs
  }
  mp := map[int64]bool{}
  for _, id := range userIDs {
    mp[id] = true
  }

  sl := make([]int64, 0, len(mp))

	for i := range mp {

    	sl = append(sl, int64(i))
	}
  
  return sl

}

func main() {
	fmt.Println(greeting.Hello(), greeting.Bye())
	color.Cyan(greeting.Bye())

	fmt.Println(UniqueUserIDs([]int64{55, 55, 33, 22}))
}
