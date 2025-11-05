package main

import "fmt"

func pessoa()(x string, n int){
	var nome string
	var idade int

	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	fmt.Printf("Olá %s, você tem %d anos.\n", nome, idade)
	return nome , idade
}
