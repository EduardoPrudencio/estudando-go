/*
	INICIALMENTE O ERRO QUE OCORRIA ERA package … is not in GOROOT
	E PARA RESOLVER BASTOU EXECUTAR O COMANDO
	go env -w GO111MODULE=off

	PARA COMPILAR O PACOTE PRECISAMOS, NO DIRETÓRIO DO PACOTE, EXECUTAR
	O COMANDO GO BUILD

	EXP: go build calculadora.go

	PARA INSTALAR UM PACOTE EXTERNO, DO GITHUB POR EXEMPLO, BASTA EXECUTAR O COMANDO
	go get -u github.com/cod3rcursos/goarea

	MESMO QUE A PASTA GITHUB.COM NÃO TENHA SIDO CRIADA ISSO VAI SER FEITO COMO PARTE DO PROCESSO.
*/

package main

import (
	"calculadora"
	"fmt"

	"github.com/cod3rcursos/goarea"
)

func main() {
	fmt.Println(calculadora.Soma(2, 1))
	fmt.Println(calculadora.Subtracao(2, 1))
	//PACOTE BAIXADO DO GITHUB
	fmt.Println(goarea.Circ(20.5))
}
