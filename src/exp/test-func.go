package main

import "fmt"

func sum(val1 int, val2 int) {
	fmt.Printf("val1: %d, val2: %d\n", val1, val2)
}

func main() {
	sum(2, 3)
	sum()
}
