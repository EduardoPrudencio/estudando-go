package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Finish!!!")
	ch <- 6
}

func ehPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func getPrimos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if ehPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c)
}

func main() {
	// ch := make(chan int, 3)
	// go routine(ch)
	// fmt.Println(<-ch)

	ch := make(chan int, 30)
	go getPrimos(cap(ch), ch)

	// O MÉTODO CAP RETORNA O TAMANHO LIMITE DO BUFFER
	//fmt.Println("***** ", cap(ch))

	for primo := range ch {
		fmt.Printf("%d -", primo)
	}
	fmt.Println("\nFim!")
}
