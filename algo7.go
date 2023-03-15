package main

import "fmt"

func main() { //Faça um algoritmo que leia um número inteiro e mostre o seu sucessor e antecessor.
	var num int
	fmt.Println("Escreva um número:")
	fmt.Scan(&num)

	ant := num - 1
	suce := num + 1
	fmt.Println("O seu antecessor é:", ant)
	fmt.Println("O seu sucessor é:", suce)
}
