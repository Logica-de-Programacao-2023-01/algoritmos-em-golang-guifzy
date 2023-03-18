package main

import (
	"fmt"
)

func main() { //Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 10% se o salário for menor que R$ 1000,00; ou de 5% se o salário for maior ou igual a R$ 1000,00.
	var salario float64
	fmt.Println("Digite seu salário:")
	fmt.Scan(&salario)

	des5 := (salario * 5) / 100
	des10 := (salario * 10) / 100

	if salario < 1000 {
		fmt.Println("O seu novo salário é de: R$", salario+des10)
	} else {
		fmt.Println("O seu novo salário é de: R$", salario+des5)
	}

}
