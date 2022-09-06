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
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func connectDb(dbName string) *sql.DB {

	fmt.Printf(" Conectando em db %s ...\n", dbName)
	serverFullName := fmt.Sprintf("root:q6dw@s460e@/%s", dbName)

	// fmt.Println(serverFullName)

	db, err := sql.Open("mysql", serverFullName)

	if err != nil {
		panic(err)
	}

	// defer db.Close()
	return db
}

func insert(db *sql.DB, userName string) {
	stmt, _ := db.Prepare("insert into users(name) values(?)")
	res, _ := stmt.Exec(userName)
	id, _ := res.LastInsertId()

	fmt.Printf("O id gerado para o insert do usuário %s foi %d\n", userName, id)
}

func update(db *sql.DB, userName string, id int64) {
	stmt, _ := db.Prepare("update users set name = ? where id = ?")
	stmt.Exec(userName, id)
}

func delete(db *sql.DB, id int64) {
	stmt, _ := db.Prepare("delete from users where id = ?")
	stmt.Exec(id)
}

func insertUserAndIdTransaction(db *sql.DB, id int64, userName string) {
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("insert into users(id, name) values(?,?)")
	_, err := stmt.Exec(id, userName)

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	} else {
		fmt.Printf("O id gerado para o insert do usuário %s foi %d\n", userName, id)
		tx.Commit()
	}
}

func getGroupById(db *sql.DB, id int64) *sql.Rows {
	rows, _ := db.Query("select id, name from users where id > ?", id)
	return rows
}

type user struct {
	id   int
	name string
}

func main() {

	dbName := "estudo_go"

	db := connectDb("")

	createTableCommand := "create database if not exists " + dbName
	exec(db, createTableCommand)
	fmt.Println("Banco de dados criado!")

	referDbCommand := "use " + dbName

	exec(db, referDbCommand)
	exec(db, "drop table if exists users")
	exec(db, `create table users (
		id integer auto_increment,
		name varchar(80),
		PRIMARY KEY (id)
	)`)

	fmt.Println("Tabela users criada!")

	db = connectDb(dbName)

	insert(db, "Maria")
	insert(db, "Anderson")
	insert(db, "Felipe")
	insert(db, "Ana")
	insert(db, "Elza")
	insert(db, "Cláudio")

	insertUserAndIdTransaction(db, 1001, "Kacashi")
	insertUserAndIdTransaction(db, 1004, "Naruto")
	insertUserAndIdTransaction(db, 1005, "Boruto")

	update(db, "Naruto Shippuden", 1004)
	delete(db, 1001)

	rows := getGroupById(db, 1)

	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name)
		fmt.Println(u)
	}

	rows.Close()
	defer db.Close()
}
