package main

import "fmt"

func main() {
	var dia int64
	fmt.Println("Digite um número inteiro entre 1 e 7")
	fmt.Scan(&dia)
	if dia == 1 {
		fmt.Println("O dia é domingo")
	} else if dia == 2 {
		fmt.Println("O dia é segunda")
	} else if dia == 3 {
		fmt.Println("O dia é terça")
	} else if dia == 4 {
		fmt.Println("O dia é quarta")
	} else if dia == 5 {
		fmt.Println("O dia é quinta")
	} else if dia == 6 {
		fmt.Println("O dia é sexta")
	} else if dia == 7 {
		fmt.Println("O dia é sábado")
	}
}
