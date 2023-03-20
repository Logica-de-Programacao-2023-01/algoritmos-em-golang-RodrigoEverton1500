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
	fmt.Println("A média ponderada é de:", ((num1*2)+(num2*3)+(num3*5))/3)
}
