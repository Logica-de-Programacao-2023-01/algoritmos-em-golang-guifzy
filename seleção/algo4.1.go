package main

import "fmt"

func main() { //Faça um algoritmo que leia a altura e o sexo de uma pessoa e mostre se ela está abaixo, dentro ou acima do peso ideal, utilizando as fórmulas do exercício 3 da lista de algoritmos sequenciais.
	var massa float64
	var altura float64
	var sexo string
	fmt.Println("Digite sua altura:")
	fmt.Scan(&altura)
	fmt.Println("Digite seu peso:")
	fmt.Scan(&massa)
	fmt.Println("Digite seu sexo(M para masculino e F para feminino):")
	fmt.Scan(&sexo)

	imc := massa / (altura * altura)
	if imc < 18 {
		fmt.Println("A baixo do peso")
	} else if sexo == "M" && sexo == "F" && imc >= 18 && imc <= 25 {
		fmt.Println("Peso ideal")
	} else if sexo == "M" && imc >= 26 && imc <= 30 {
		fmt.Println("Sobrepeso")
	} else if sexo == "F" && imc >= 18 && imc <= 24 {
		fmt.Println("Peso ideal")
	} else {
		fmt.Println("Sobrepso")
	}

}
