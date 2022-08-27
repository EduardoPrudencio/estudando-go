package main

import "fmt"

func getValues() (bool, int) {
	return true, 1000
}

func reverseValues(valueOne int, valueTwo int) (firstValue int, secondValue int) {
	firstValue = valueTwo
	secondValue = valueOne
	return
}

func mult(a int, b int) int {
	return a * b
}

func exec(function func(a int, b int) int, v1 int, v2 int) int {
	return function(v1, v2)
}

func media(numbers ...int) int {

	var total int = 0

	for _, num := range numbers {
		total += num
	}

	return total / len(numbers)
}

func main() {

	success, count := getValues()
	fmt.Printf("O resultado da verificação foi %t com %d unidades.\n", success, count)

	/*
		PARA IGNORAR UM DOS VALORES EM FUNÇÕES QUE RETORNAL MAIS DE UM VALOR DEVEMOS USAR _ NO LUGAR DESSE
		VALOR
	*/

	_, countSecond := getValues()
	fmt.Printf("Este exemplo retorna %d unidades.\n", countSecond)

	tirdSucess, _ := getValues()
	fmt.Printf("Este exemplo retorna %t.\n", tirdSucess)

	/*
		RETORNO LIMPO
		GO PERMITE QUE AS VARIÁVEIS DE RETORNO SEJAM DEFINIDAS NA ASSINATURA
		DA FUNÇÃO.

		EX: func getTwoValues() (firstValue int, secondValue int)

		BASEADO NESTE EXEMPLO TEREMOS COMO RETORNO DUAS VARIÁVEIS DO TIPO INT
	*/

	// func reverseValues(valueOne int, valueTwo int) (firstValue int, secondValue int) {
	// 	firstValue = valueTwo
	// 	secondValue = valueOne
	// 	return
	// }

	a, b := reverseValues(1, 2)
	fmt.Println(a, b)

	/*
		ARMAZENANDO FUNÇÕES EM VARIÁVEIS
	*/

	sub := func(a int, b int) int {
		return a - b
	}

	fmt.Println(sub(3, 4))

	/*
		RECEBENDO FUNÇÕES COMO PARÂMETROS
	*/

	fmt.Println(exec(mult, 2, 5))

	/*
		FUNÇÕES VARIÁTICAS
	*/


	valores := []int{10,20,40,30}
	fmt.Println("A média desses números é: ",media(valores...))

}
