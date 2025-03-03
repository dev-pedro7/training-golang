package main

import "fmt"

func sumTwoNumbers(prm01 int, prm02 int) int {
	return prm01 + prm02
}

func sumAll(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	result := sumTwoNumbers(19551, 15158)
	fmt.Println("Resultado:", result)
	fmt.Println(sumAll(10, 10, 10, 10))
}
