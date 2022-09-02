package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
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

func fastResponse(url1, url2, url3 string) string {
	c1 := getTitle(url1)
	c2 := getTitle(url2)
	c3 := getTitle(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Timeout!"
		// default:
		// 	return "Sem respostas!"
	}
}

func talk(sender string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			channel <- fmt.Sprintf("%s falando: %d", sender, i)
		}
	}()
	return channel
}

func joinTalk(chanelOne, channelTwo <-chan string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			select {
			case text := <-chanelOne:
				channel <- text
			case text := <-channelTwo:
				channel <- text
			}
		}
	}()
	return channel
}

func main() {
	// t1 := getTitle("https://google.com.br", "https://youtube.com")
	// fmt.Println(<-t1)
	// fmt.Println(<-t1)

	// c := join(
	// 	getTitle("https://google.com.br", "https://youtube.com"),
	// 	getTitle("https://www.cod3r.com.br", "https://www.amazon.com"),
	// )

	// fmt.Println(<-c, "|", <-c)
	// fmt.Println(<-c, "|", <-c)

	champion := fastResponse(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
	)

	fmt.Println(champion)

	fmt.Println("************************************")
	fmt.Println("   Executando o método joinTalk")
	fmt.Println("************************************")

	talkChannel := joinTalk(talk("Anna"), talk("Pedro"))

	fmt.Println(<-talkChannel, <-talkChannel)
	fmt.Println(<-talkChannel, <-talkChannel)
	fmt.Println(<-talkChannel, <-talkChannel)
}
