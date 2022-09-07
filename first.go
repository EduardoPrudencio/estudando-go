package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func soma(numeroUm int, numeroDois int) int {
	return numeroUm + numeroDois
}

func parOuImpar(numero int) {
	if numero%2 == 0 {
		fmt.Printf("O número %d é par", numero)
	} else {
		fmt.Printf("O número %d é impar", numero)
	}
}

func meDizSeutipo(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Inteiro")
	case float32, float64:
		fmt.Println("Ponto flutuante")
	case func():
		fmt.Println("Função")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Não sei dizer")
	}
}

func main() {
	/*
		A linguagem Go é fortemente tipada e as variáveis podem ser criadas com seus tipos
		definidos de forma explicita ou implicita.

		var valorDeTeste int = 2 (Definição explicita)
		valorDeTeste := 2 (Inferência de valor)
	*/

	nome := ""
	sobreNome := ""
	fmt.Println("Informe um nome e sobrenome")
	fmt.Scanln(&nome, &sobreNome)
	fmt.Println("O nome e sobrenome informado pelo usuário foram:git ", nome, sobreNome)

	const valorDeConstante int = 2
	var valorDeVariavel int = 2
	numeroRandomico := rand.Intn(1000)

	variavel1 := 1
	variavel2 := true
	variavel3 := "Texto"
	variavel4 := time.Now()

	fmt.Println("Exibindo tipos de dados criados.")
	fmt.Println("O valor da constante valorDeTeste :", valorDeConstante)
	fmt.Println("O valor da variável valorDeTeste :", valorDeVariavel)
	fmt.Println("O valor da variável numeroRandomico :", numeroRandomico)

	fmt.Println("O tipo da variável variavel1 :", reflect.TypeOf(variavel1))
	fmt.Println("O tipo da variável variavel2 :", reflect.TypeOf(variavel2))
	fmt.Println("O tipo da variável variavel3 :", reflect.TypeOf(variavel3))
	fmt.Println("O tipo da variável variavel4 :", reflect.TypeOf(variavel4))

	fmt.Println("Exibindo o resultado de um método que soma dois números, nesse exemplo, 10 e 25.", soma(10, 25))

	fmt.Println("\n\n***************")
	fmt.Println("   PONTEIROS")
	fmt.Println("***************")

	var primeiroValor int
	var enderecoDoPrimeiroValor *int

	primeiroValor = 10
	enderecoDoPrimeiroValor = &primeiroValor

	fmt.Println("O valor da variável primeiroValor é:", primeiroValor)
	fmt.Println("O endereço da variável primeiroValor é:", enderecoDoPrimeiroValor)

	primeiroValor++
	fmt.Println("Estamos somando 1 à variável primeiroValor e agora o seu resultado é:", primeiroValor)

	*enderecoDoPrimeiroValor++
	fmt.Println("Estamos somando 1 ao endereço enderecoDoPrimeiroValor e agora o valor da variável primeiroValor é:", primeiroValor)

	fmt.Println("\n\n***********************")
	fmt.Println("   ESTRUTURA DE CONTROLE")
	fmt.Println("***************************")

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	numeroDinamico := random.Intn(10)
	parOuImpar(int(numeroDinamico))

	if i := random.Intn(10); i > 5 {
		fmt.Printf("\nO número %d é maior que 5", i)
	} else {
		fmt.Printf("\nO número %d é menor que 6", i)
	}

	/*
		NÃO EXISTE WHILE EM GO!!!!!
	*/

	fmt.Println("\n\n***************************")
	fmt.Println("   USANDO FOR COMO WHILE")
	fmt.Println("***************************")

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\n\n**************************************")
	fmt.Println("   USANDO FOR DE FORMA CONVENCIONAL")
	fmt.Println("**************************************")

	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	/*
			O COMANDO FOR PODE EXECUTAR UM BLOCO INFINITAMENTE SEGUINDO O EXEMPLO:

			for {
				fmt.Println("Teste")
				time.Sleep(time.Second *3)
		    }
	*/

	fmt.Println("\n\n********************")
	fmt.Println("   COMANDO SWITCH")
	fmt.Println("********************")

	/*
		ESSE BLOCO NÃO FOI REPETIDO PORQUE SUA DEFINIÇÃO ACONTECE ANTES
			source := rand.NewSource(time.Now().UnixNano())
			random := rand.New(source)
	*/
	var valor int = random.Intn(10)

	switch valor {
	case 1, 2:
		fmt.Println("E")
	case 3, 4:
		fmt.Println("D")
	case 5, 6:
		fmt.Println("C")
	case 7, 8:
		fmt.Println("B")
	case 9, 10:
		fmt.Println("A")

	default:
		fmt.Println("Valor não reconhecido")
	}

	meDizSeutipo(22)
	meDizSeutipo(22.789)
	meDizSeutipo(false)
	meDizSeutipo(func() {})
	meDizSeutipo("Teste")

	fmt.Println("\n\n*************************")
	fmt.Println("   ARRAY, SLICES E MAP")
	fmt.Println("*************************")

	integerList := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for i, numero := range integerList {
		fmt.Printf("%d _ %d\n", i+1, numero)
	}

	fmt.Println()

	/*
		SLICE
	*/

	slice1 := integerList[0:5]

	fmt.Println(slice1)

	/*
		MAP
	*/

	clientes := make(map[int]string)

	clientes[258742] = "João"
	clientes[298325] = "Maria"
	clientes[211223] = "Fernanda"
	clientes[232312] = "Francisco"

	//fmt.Println(clientes)
	delete(clientes, 211223)

	for index, name := range clientes {
		fmt.Printf("%s - %d\n", name, index)
	}

	funcionariosESalarios := map[string]float64{
		"Antunes": 2567.45,
		"Waldeir": 2567.45,
		"Walter":  5587.15,
	}

	//fmt.Println(funcionariosESalarios)
	for index, name := range funcionariosESalarios {
		fmt.Printf("O funcionário %f ganha %s por mês\n", name, index)
	}

	/*
		MAP ALINHADO
	*/

	listaDeALunosPorFaixa := map[string]map[int]string{
		"Branca": {
			1: "Fernando Aguiar",
			2: "Andreson Pereira",
		},
		"Azul": {
			3: "Robson Nunes",
			4: "André Silva",
		},
		"Roxa": {
			5: "Raider Costa",
			6: "Mário Rodovário",
		},
	}

	for faixa, aluno := range listaDeALunosPorFaixa {
		//fmt.Printf("Faixa %s\n %v", faixa, aluno)
		fmt.Println(faixa, aluno)
	}
}
