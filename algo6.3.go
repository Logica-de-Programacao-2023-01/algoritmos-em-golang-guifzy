package main

import "fmt"

func main() { //Faça um algoritmo que imprima a tabuada de multiplicação de 1 a 10 para um número fornecido pelo usuário.
	var num int
	fmt.Println("Digite um número:")
	fmt.Scan(&num)

	for mult := 1; mult <= 10; mult++ {
		fmt.Println("O resulado da tabuada é:", num*mult)
	}
}
