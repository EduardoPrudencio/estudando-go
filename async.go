package main

import (
	"fmt"
	"runtime"
	"time"
)

func taskFake(numTask int, count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("Task %d - executando %d vez(es).", numTask, i)
	}
}

func sendChannel(value int, c chan int) {
	c <- value
	time.Sleep(time.Second)
	c <- value * 2
	time.Sleep(time.Second)
	c <- value * 3
	time.Sleep(time.Second)
}

func main() {
	fmt.Printf("Temos %d cores nesse computador.\n\n", runtime.NumCPU())

	// A PALAVRA GO PERMITE QUE O MÉTODO SEJA EXECUTADO DE FORMA CONCORRENTE MAS
	// SE A THREAD PRINCIPAL NÃO ESTIVER EXECUTANDO AS THREADS CONCORRENTES SERÃO FINALIZADAS
	//go taskFake(1, 100)
	//taskFake(2, 10)

	//GO PERMITE A CRIAÇÃO DE CANAIS PARA FACILITAR A COMUNICAÇÃO
	//ch1 := make(chan int, 1)
	//ch1 <- 1487

	//fmt.Println(<-ch1)

	ch2 := make(chan int)
	go sendChannel(1, ch2)

	// a, b := <-ch2, <-ch2
	// fmt.Println(a, b)

	// c := <-ch2
	// fmt.Println(c)

	i := 1
	for i <= 3 {
		fmt.Println(<-ch2)
		i++
	}

}
