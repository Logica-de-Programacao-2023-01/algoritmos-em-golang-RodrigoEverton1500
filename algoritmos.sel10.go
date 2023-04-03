package main

import "fmt"

func main() {
	var idade int64
	fmt.Println("Digite a sua idade")
	fmt.Scan(&idade)
	if idade <= 9 {
		fmt.Println("Sua idade é mirim")
	} else if idade >= 10 && idade <= 13 {
		fmt.Println("Sua idade é infatil")
	} else if idade >= 13 && idade <= 17 {
		fmt.Println("Sua idade é juvenil")
	} else if idade >= 18 {
		fmt.Println("Sua idade é adulto")
	}
}
