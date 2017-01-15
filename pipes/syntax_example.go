package main

import "fmt"

func a() []string {
	return []string{"a"}
}

func empty(a interface{}) bool {
	return true
}

func count(a []string) int {
	return len(a)
}

func betterThan12(n int) bool {
	return n != 12
}

func main() {
	a()
	:> empty
	?> fmt.Println("ok")
		:> fmt.Println("oh yeah, that's good")
	!?> fmt.Println("not ok")
		:> fmt.Println("damn, it wasn't empty")
	:> count
	:> betterThan12
	?> fmt.Println("better than 12")
	!?> fmt.Println("nope")
}
