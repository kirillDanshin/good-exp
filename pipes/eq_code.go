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
	tmpA := a()
	if empty(tmpA) {
		fmt.Println("ok")
		fmt.Println("oh yeah, that's good")
	} else {
		fmt.Println("not ok")
		fmt.Println("damn, it wasn't empty")
	}
	if betterThan12(count(tmpA)) {
		fmt.Println("better than 12")
	} else {
		fmt.Println("nope")
	}
}
