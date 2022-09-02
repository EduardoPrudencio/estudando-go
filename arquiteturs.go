package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOARCH) // EXIBE ARQUITETURA
	fmt.Println(runtime.GOOS)   // EXIBE SISTEMA OPERACIONAL
}
