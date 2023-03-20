package main

import "fmt"

func main() {

	var preço int64
	fmt.Println("Digite o preço do produto")
	fmt.Scan(&preço)
	fmt.Println("O preço com desconto é:", (preço/10)*9)
}
