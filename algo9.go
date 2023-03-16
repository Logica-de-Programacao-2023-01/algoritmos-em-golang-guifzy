package main

import "fmt"

func main() { //Faça um algoritmo que leia o preço de um produto e mostre o seu valor com desconto de 10%.
	var preco float64
	fmt.Println("Digite o preço do produto: R$")
	fmt.Scan(&preco)

	desconto := preco * 10 / 100
	fmt.Print("O desconto do produtoo é de: R$", desconto)

}
