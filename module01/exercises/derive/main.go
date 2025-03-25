package main

import "fmt"

func Derive(coefficient, exponent int) string {
	return fmt.Sprintf("%dx^%d", coefficient*exponent, exponent-1)
}

func main() {
	fmt.Println(Derive(5, 10))
	fmt.Println(Derive(7, 3))
	fmt.Println(Derive(2, 4))

}
