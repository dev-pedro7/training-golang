package main

import "fmt"

func inline(cond bool, v1, v2 string) string {
	if cond {
		return v1
	}
	return v2
}

func main() {
	num := 777
	result := inline(num%2 == 0, "Even", "Odd")
	fmt.Println(result)
}
