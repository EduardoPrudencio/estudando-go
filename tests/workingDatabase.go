/*
	INSTALANDO O DRIVER PARA INTEGRAÇÃO COM MYSQL
	DOCKER DE CRIAÇÃO DO BANCO DE DADOS MYSQL
	docker run --name mysql-db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=q6dw@s460e -d mysql
	go get -u github.com/go-sql-driver/mysql
*/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	fmt.Println(" Conectando...")
	db, err := sql.Open("mysql", "root:q6dw@s460e@/")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "create database if not exists cursogo")
	fmt.Println("Tudo ok!")

}
