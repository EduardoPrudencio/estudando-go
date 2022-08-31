/*
	INICIALMENTE O ERRO QUE OCORRIA ERA package … is not in GOROOT
	E PARA RESOLVER BASTOU EXECUTAR O COMANDO
	go env -w GO111MODULE=off
*/

package main

import (
	"calculadora"
	"fmt"
)

func main() {
	fmt.Println(calculadora.Soma(2, 1))
}
