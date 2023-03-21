package main

import "fmt"

func main() { //Faça um algoritmo que leia dois números inteiros e mostre o resultado da multiplicação entre eles, se ambos forem positivos; ou a soma, se pelo menos um deles for negativo.
	var a, b int
	fmt.Println("Digite dois números:")
	fmt.Scan(&a, &b)

	if a >= 0 && b >= 0 {
		fmt.Println(a * b)
	} else if a < 0 || b < 0 {
		fmt.Println(a + (b))

	}
}
