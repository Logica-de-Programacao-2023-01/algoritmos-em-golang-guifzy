package main

import "fmt"

func main() { //Faça um algoritmo que leia a idade de uma pessoa e mostre a sua classificação, de acordo com a seguinte tabela:
	var idade int
	fmt.Println("Digite a sua idade:")
	fmt.Scan(&idade)
	if idade <= 9 {
		fmt.Println("mirim")

	} else if idade >= 10 && idade <= 13 {
		fmt.Println("infantil")

	} else if idade >= 14 && idade <= 17 {
		fmt.Println("juvenil")

	} else {
		fmt.Println("adulto")
	}
}
