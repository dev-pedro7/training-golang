package main

import "fmt"

func main() {
	var hour int
	fmt.Print("Digite a Hora: ")
	fmt.Scanf("%d", &hour)

	switch {
	case hour <= 12:
		fmt.Printf("Bom dia, são %d horas.\n", hour)
	case hour >= 13 && hour < 19:
		fmt.Printf("Boa tarde, são %d horas.\n", hour)
	default:
		fmt.Printf("Boa noite, são %d horas.\n", hour)
	}
}
