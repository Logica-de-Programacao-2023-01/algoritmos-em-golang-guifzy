package main

import "fmt"

func main() { //Faça um algoritmo que leia o número de dias trabalhados por um funcionário e o valor da sua diária e calcule o seu salário.
	var dias float64
	var diaria float64
	fmt.Println("Digite a quantidade de dias de trabalho:")
	fmt.Scan(&dias)
	fmt.Print("Digite o valor da diária: R$")
	fmt.Scan(&diaria)

	salario := diaria * dias
	fmt.Print("Seu salário é de: R$", salario)

}
