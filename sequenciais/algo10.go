package main

import "fmt"

func main() { //Faça um algoritmo que leia o peso de uma pessoa em quilos e converta para libras.
	var quilos float64
	fmt.Println("Digite o valor em quilos:")
	fmt.Scan(&quilos)

	libras := quilos * 2.2046
	fmt.Println("A coversão do peso em libras é de:", libras)

}
