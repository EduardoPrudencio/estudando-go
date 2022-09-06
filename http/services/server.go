package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/user/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		fmt.Println(w, "O id é :", id)
		userById(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry!...:")
	}
}

func userById(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:q6dw@s460e@/estudo_go")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var user User
	db.QueryRow("select id, name from users where id =?", id).Scan(&user.Id, &user.Name)
	jsonResponse, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonResponse))
}

func main() {
	http.HandleFunc("/user/", GetUserById)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
