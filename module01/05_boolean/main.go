package main

import "fmt"

func main() {
	var var1 bool = true
	var2, var3 := false, true

	fmt.Println("AND: ", var1 && var2)
	fmt.Println("OR: ", var2 || var3)
	fmt.Println("NOT: ", var1)

	varnum1, varnum2 := 100, 200
	varnum3 := varnum1 > varnum2
	fmt.Printf("%v é maior que %v?:%v\n", varnum1, varnum2, varnum3)
}
