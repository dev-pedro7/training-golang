package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Lê
	data, err := os.ReadFile("../create/ArquivoCriado.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Exibe
	fmt.Println(string(data))
}
