package main

import "fmt"

func main() { //Faça um algoritmo que leia um número inteiro e mostre se ele é múltiplo de 3 e de 5 ao mesmo tempo.
	var num int
	fmt.Println("Digite um número:")
	fmt.Scan(&num)

	if num%3 == 0 {
		fmt.Println("É múltiplo de 3")
	} else if num%5 == 0 {
		fmt.Println("É múltiplo de 5")
	} else {
		fmt.Println("Não é múltiplo de 3 nem de 5.")
	}
}
