package main

import (
	"fmt"
	"scratchs/leetcode/design_browser_history"
)

func main() {
	homepage := "google"

	obj := design_browser_history.Constructor(homepage)

	obj.Visit("google+1")
	obj.Visit("google+2")
	obj.Visit("google+3")

	prev := obj.Back(2)
	next := obj.Forward(1)

	fmt.Println(prev, next)
}
