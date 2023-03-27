package main

import "fmt"

func main() {
	var num1 int64

	fmt.Println("Digite um número:")
	fmt.Scan(&num1)
	
	if num1%3 == 0 {
		fmt.Println("O número é múltiplo de 3")
	} else {
		fmt.Println("O número não é múltiplo 3")
	}
	if num1%5 == 0 {
		fmt.Print("O número é múltiplo de 5")
	} else {
		fmt.Println("O número não é múltiplo de 5")
	}

}
