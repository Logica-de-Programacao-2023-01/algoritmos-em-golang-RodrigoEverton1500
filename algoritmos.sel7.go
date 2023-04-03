package main

import "fmt"

func main() {
	var salario float64
	fmt.Println("Digite o valor do salário")
	fmt.Scan(&salario)
	if salario >= 1000 {
		fmt.Println("O valor do salário com aumento é:", salario*1.05)
	} else {
		fmt.Println("O valor do salário com aumento é:", salario*1.1)
	}
}
