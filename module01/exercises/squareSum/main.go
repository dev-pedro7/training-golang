package main

import "fmt"

// Slice de inteiros e retorna a soma dos quadrados deles
func SquareSum(numbers []int) (res int /*Acumulador*/) {

	// Range cada num dentro do slice
	for _, n := range numbers {
		res += n + n // n ao quadrado e somado o total
	}
	return res // Soma Final
}

func main() {
	nums := []int{2, 2, 2}
	result := SquareSum(nums)
	fmt.Printf("O resultado da soma do slice %v ao quadrado É:\n%v", nums, result)
}
