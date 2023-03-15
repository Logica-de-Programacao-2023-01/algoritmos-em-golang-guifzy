package main

import "fmt"

func main() { //Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo.
	var a int
	fmt.Println("Escreva um número:")
	fmt.Scan(&a)

	dobro := a * 2
	triplo := a * 3
	quadruplo := a * 4
	fmt.Println("O dobro é:", dobro)
	fmt.Println("O triploo é:", triplo)
	fmt.Println("O quadruplo é:", quadruplo)

}
