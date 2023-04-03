package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var num3 float64
	fmt.Println("Digite o primeiro número")
	fmt.Scan(&num1)
	fmt.Println("Digite o segundo número")
	fmt.Scan(&num2)
	fmt.Println("Digite o terceiro número")
	fmt.Scan(&num3)
	if num1 > (num2+num3)/2 {
		fmt.Println("O primeiro número é o maior")
	} else if num2 > (num1+num3)/2 {
		fmt.Println("O segundo número é o maior")
	} else if num3 > (num2+num1)/2 {
		fmt.Println("O terceiro número é o maior")
	}
}
