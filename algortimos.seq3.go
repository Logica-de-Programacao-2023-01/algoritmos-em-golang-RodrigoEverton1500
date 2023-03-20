package main

import "fmt"

func main() {

	var peso float64
	var altura float64
	fmt.Println("Digite o seu peso:")
	fmt.Scan(&peso)
	fmt.Println("Digite a sua altura")
	fmt.Scan(&altura)
	fmt.Println("Seu IMC é:", peso/altura)
}
