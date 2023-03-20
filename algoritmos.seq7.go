package main

import "fmt"

func main() {

	var num1 int64
	fmt.Println("Digite um número inteiro:")
	fmt.Scan(&num1)
	fmt.Println("Seu antecessor é:", num1-1)
	fmt.Println("Seu sucessor é:", num1+1)

}
