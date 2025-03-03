package main

import "fmt"

func main() {
	var hour int
	fmt.Print("Digite a Hora: ")
	fmt.Scanf("%v", &hour)
	if hour <= 12 {
		fmt.Printf("Bom dia, são %d Horas", hour)
	} else if hour >= 13 && hour < 19 {
		fmt.Printf("Boa tarde, são %d Horas", hour)
	} else {
		fmt.Printf("Boa noite, são %d Horas", hour)
	}
}
