package main

import "fmt"

func main() {

	var num1 int64
	var num2 int64
	fmt.Println("Digite um número:")
	fmt.Scan(&num1)
	fmt.Println("Digite outro número:")
	fmt.Scan(&num2)
	if num1 > num2 {
		fmt.Println(num1, "é maior que", num2)
	} else if num2 > num1 {
		fmt.Println(num2, "é maior que", num1)
	} else {
		fmt.Println("Os dois números são iguais")
	}

}
