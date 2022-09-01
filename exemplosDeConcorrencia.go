package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
	MÉTODO QUE RETORNA UM CANAL DE SOMENTE LEITURA
*/

func getTitle(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			title, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- title.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func sendMessageToChannel(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func join(entrada1 <-chan string, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go sendMessageToChannel(entrada1, c)
	go sendMessageToChannel(entrada2, c)
	return c
}

func main() {
	// t1 := getTitle("https://google.com.br", "https://youtube.com")
	// fmt.Println(<-t1)
	// fmt.Println(<-t1)

	c := join(
		getTitle("https://google.com.br", "https://youtube.com"),
		getTitle("https://www.cod3r.com.br", "https://www.amazon.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
