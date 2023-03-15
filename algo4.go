package main

import "fmt"

func main() { //Faça um algoritmo que leia três números reais e mostre a média ponderada entre eles, com pesos 2, 3 e 5, respectivamente.
	var num1 float64
	var num2 float64
	var num3 float64
	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&num1)
	fmt.Println("Digite o segundo número:")
	fmt.Scan(&num2)
	fmt.Println("Digite o terceiro número:")
	fmt.Scan(&num3)

	peso2 := (num1*2 + num2*2 + num3*2) / 10
	peso3 := (num1*3 + num2*3 + num3*3) / 10
	peso5 := (num1*5 + num2*5 + num3*5) / 10
	fmt.Println("A média com peso 2 é:", peso2)
	fmt.Println("A média com peso 3 é:", peso3)
	fmt.Println("A média com peso 5 é:", peso5)

}
