package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func curretHour(w http.ResponseWriter, r *http.Request) {
	fullHour := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Current Hour: %s</h1>", fullHour)
}

func main() {
	http.HandleFunc("/currenthour", curretHour)
	log.Println("Executando")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
