package main

import "fmt"

func main() {
	var decNumber int
	fmt.Print("Decimal para Binario: ")
	fmt.Scanf("%d", &decNumber)
	fmt.Printf("%b", decNumber)
}

// .exe em GO: go build -o program.exe main.go
