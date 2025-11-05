package main

import "fmt"
/*
3️⃣ Crie uma função ehPar(n int) que:

retorna true se o número for par;

imprime se o número é par ou ímpar no main().
*/
func main()  {



	resultado1 := ehPar(5)
	resultado2 := ehPar(2)

	fmt.Println("Resultado 1 :", resultado1)
	fmt.Println("Resultado 2 :", resultado2)

}

func ehPar(n int) bool{
		if (n % 2) == 0 {
			return true
		}
		return false
}
