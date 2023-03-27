package main

import "fmt"

func main() {

	var peso float64
	fmt.Println("Qual é o seu peso em quilos?")
	fmt.Scan(&peso)
	fmt.Println("Seu peso em libras é:", peso*2.2)

}
