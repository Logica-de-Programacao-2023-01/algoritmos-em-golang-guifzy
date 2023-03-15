package main

import "fmt"

func main() { //Faça um algoritmo que leia o peso e a altura de uma pessoa e calcule o seu IMC (Índice de Massa Corporal).
	var massa float64
	var altura float64
	fmt.Println("Digite sua altura:")
	fmt.Scan(&altura)
	fmt.Println("Digite seu peso:")
	fmt.Scan(&massa)

	imc := massa / (altura * altura)
	fmt.Println("O seu IMC é:", imc)
}
