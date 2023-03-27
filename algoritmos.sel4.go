package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	var sexo int64
	fmt.Println("Digite o seu peso:")
	fmt.Scan(&peso)
	fmt.Println("Digite a sua altura:")
	fmt.Scan(&altura)
	fmt.Println("Digite o seu sexo")
	fmt.Println("Digite 1 para masculino e 2 para feminino:")
	fmt.Scan(&sexo)

	if sexo == 1 {
		fmt.Println(peso / (altura * altura))
	}

	if (peso / (altura * altura)) > 24.9 {
		fmt.Println("Você é obeso")
	} else if (peso / (altura * altura)) < 20 {
		fmt.Println("Você está magro")
	} else {
		fmt.Println("Você está dentro do peso ideal")
	}

	if sexo == 2 {
		fmt.Println(peso / (altura * altura))
	}

	if (peso / (altura * altura)) > 24.9 {
		fmt.Println("Você é obeso")
	} else if (peso / (altura * altura)) < 18.5 {
		fmt.Println("Você está magro")
	} else {
		fmt.Println("Você está dentro do peso ideal")
	}
}
