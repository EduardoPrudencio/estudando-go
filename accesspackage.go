/*
	INICIALMENTE O ERRO QUE OCORRIA ERA package … is not in GOROOT
	E PARA RESOLVER BASTOU EXECUTAR O COMANDO
	go env -w GO111MODULE=off

	PARA COMPILAR O PACOTE PRECISAMOS, NO DIRETÓRIO DO PACOTE, EXECUTAR
	O COMANDO GO BUILD

	EXP: go build calculadora.go
*/

package main

import (
	"calculadora"
	"fmt"
)

func main() {
	fmt.Println(calculadora.Soma(2, 1))
	fmt.Println(calculadora.Subtracao(2, 1))
}
