package main

import (
	"log"
	"os"
)

func main() {
	// os.Create para criar o arquivo, informando nome e extensão
	file, err := os.Create("ArquivoCriado.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Escreve no "ArquivoCriado.txt"
	_, err = file.WriteString("Olá Mundo! \n gerado em Go.")
	if err != nil {
		log.Fatal(err)
	}
}
