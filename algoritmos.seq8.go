package main

import "fmt"

func main() {

	var dias float64
	var diaria float64
	fmt.Println("Digite o número de dias trabalhados")
	fmt.Scan(&dias)
	fmt.Println("Digite o valor da diária")
	fmt.Scan(&diaria)
	fmt.Println("O salário é de:", dias*diaria)
}
