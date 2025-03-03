package main

import "fmt"

func trimestre(month int) (result int) {
	switch month {
	case 1, 2, 3:
		result = 1
	case 4, 5, 6:
		result = 2
	case 7, 8, 9:
		result = 3
	case 10, 11, 12:
		result = 4
	}
	return
}

func main() {
	var input int
	fmt.Printf("Digite um mês:")
	fmt.Scanf("%d", &input)
	fmt.Printf("O Trimestre do mês %d é o: %v", input, trimestre(input))
}
