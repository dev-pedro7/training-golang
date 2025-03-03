package main

import "fmt"

func max(x int) int {
	return 100 + x
}

func main() {
	result := max(50)
	fmt.Println(result)
}
