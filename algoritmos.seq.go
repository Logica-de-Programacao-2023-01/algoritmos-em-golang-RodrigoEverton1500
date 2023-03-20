package main

import "fmt"

func main() {

	var num1 int64
	var num2 int64
	var num3 int64

	fmt.Println("Digite o primeiro número")
	fmt.Scan(&num1)
	fmt.Println("Digite o segundo número")
	fmt.Scan(&num2)
	fmt.Println("Digite o terceiro número")
	fmt.Scan(&num3)
	fmt.Println(num2 + num1 + num3)
}
