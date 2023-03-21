package main

import "fmt"

func main() { //Faça um algoritmo que leia três números inteiros e mostre o menor deles.
	var a, b, c int
	fmt.Println("Digite 3 números inteiros:")
	fmt.Scan(&a, &b, &c)

	if a <= b && a <= c {
		fmt.Println("O menor número é:", a)

	} else if b <= a && b <= c {
		fmt.Println("O menor número é:", b)

	} else {
		fmt.Println("O menor número é:", c)

	}

}
