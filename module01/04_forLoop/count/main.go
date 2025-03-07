package main

import "fmt"

func main() {
	for p := 0; p <= 5; p++ {
		fmt.Println("Contando:", p)
	}

	c := 4

	for c >= 0 {
		fmt.Println("Contando:", c)
		c--
	}

}
