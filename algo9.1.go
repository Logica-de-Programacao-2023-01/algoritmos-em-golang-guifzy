package main

import "fmt"

func main() { //Faça um algoritmo que leia três números reais e mostre-os em ordem crescente.
	var num1, num2, num3 int
	fmt.Println("Digite 3 números inteiros:")
	fmt.Scan(&num1, &num2, &num3)

	if num1 <= num2 && num1 <= num3 {
		fmt.Print(num1, " ")
		if num2 <= num3 {
			fmt.Print(num2, " ", num3)
		} else {
			fmt.Print(num3, " ", num2)
		}
	} else if num2 <= num1 && num2 <= num3 {
		fmt.Print(num2, " ")
		if num1 <= num3 {
			fmt.Print(num1, " ", num3)
		} else {
			fmt.Print(num3, " ", num1)
		}
	} else {
		fmt.Print(num3, " ")
		if num1 <= num2 {
			fmt.Print(num1, " ", num2)
		} else {
			fmt.Print(num2, " ", num1)
		}
	}
}

