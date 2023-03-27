package main

import "fmt"

func main() {
	var num1 int64

	fmt.Println("Digite um número:")
	fmt.Scan(&num1)

	if num1%2 == 0 {
		fmt.Println(num1, "é um número par")
	} else {
		fmt.Println(num1, "é um número impar")
	}
}
