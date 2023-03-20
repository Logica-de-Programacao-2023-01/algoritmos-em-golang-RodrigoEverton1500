package main

import "fmt"

func main() {

	var salario float64
	fmt.Println("Digite salário")
	fmt.Scan(&salario)
	fmt.Println("Salário com aumento é igual a:", (salario/10)*11.5)

}
