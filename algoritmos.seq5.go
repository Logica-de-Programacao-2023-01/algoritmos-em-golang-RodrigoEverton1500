package main

import "fmt"

func main() {

	var idade int64
	var idade2 int64
	fmt.Println("Digite a sua idade")
	fmt.Scan(&idade)
	idade2 = idade / 4
	idade = idade*365 + idade2
	fmt.Println("Sua idade é de:", idade, "dias")

}
