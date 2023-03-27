package main

import "fmt"

func main() {

	var num1 int64
	var num2 int64
	var num3 int64
	fmt.Println("Digite um número:")
	fmt.Scan(&num1)
	fmt.Println("Digite um segundo número:")
	fmt.Scan(&num2)
	fmt.Println("Digite um terceiro número:")
	fmt.Scan(&num3)
	if num1 < (num2+num3)/2 {
		fmt.Println(num1, "é o menor número")
	} else if num2 < (num1+num3)/2 {
		fmt.Println(num2, "é o menor número")
	} else {
		fmt.Println(num3, "é o menor número")
	}

}
