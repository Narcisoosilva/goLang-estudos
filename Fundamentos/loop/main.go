package  main

import	"fmt"


/*
2️⃣ Crie um programa que:

conte de 1 até 10 com um for;

pule o número 5 (use continue);

pare o loop no número 8 (use break).

*/
/*
func main()  {

	for i := 1; i < 10; i++ {
		if i == 5{
			continue
		}
		if i == 8{
			break
		}
		fmt.Printf("%d, ", i,)

	}
}
*/
/*
🧩 Nível 1 – Aquecimento (variáveis, laços e condições)

1️⃣ Verificação de idade
Crie um programa que leia a idade do usuário e:

diga se ele é menor de idade, adulto ou idoso.
(use if, else if, else)
*/

func main()  {
	var nome string
	var idade int

	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	if idade >= 60 {
		fmt.Printf("Olá %s, você tem %d anos, e já é Idoso\n", nome, idade)

	}
	if idade <= 17  {
		fmt.Printf("Olá %s, você tem %d anos, e você é menor de idade.\n", nome, idade)
	}
	if (idade >= 18)&&(idade < 60) {
		fmt.Printf("Olá %s, você tem %d anos, e você é adulto.\n", nome, idade)
	}

}
