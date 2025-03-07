package main

import "fmt"

func main() {
	arr := [3]int{100, 200, 300}
	// fmt.Println(arr[2])

	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
