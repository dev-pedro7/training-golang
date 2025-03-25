package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30}

	for i, n := range numbers {
		fmt.Printf("Indice: %v\nValor: %v\n", i, n)
	}
}
