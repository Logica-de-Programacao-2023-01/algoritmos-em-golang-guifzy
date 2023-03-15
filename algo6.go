package main

import "fmt"

func main() { //Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 15%.
	var salario float64
	fmt.Print("Digite o valor do seu salário: R$")
	fmt.Scan(&salario)

	aumento := (salario * 15) / 100
	fmt.Println("O seu aumento é de: R$", aumento)
}
