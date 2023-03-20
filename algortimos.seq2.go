package main

import "fmt"

func main() {

	var num1 int64
	fmt.Println("Digite um número")
	fmt.Scan(&num1)
	fmt.Println("O seu dobro é:", num1*2)
	fmt.Println("O seu triplo é:", num1*3)
	fmt.Println("O seu quadruplo é:", num1*4)
}
