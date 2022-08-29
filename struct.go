package main

import "fmt"

type shower interface {
	toString() string
}

type person struct {
	id       int
	name     string
	lastName string
}

/*
	COMPOSIÇÃO/HERANÇA
*/

type student struct {
	person
	note string
}

type teacher struct {
	person
}

func (p person) toString() string {
	return p.name + " " + p.lastName
}

func (p person) getFullName() string {
	return p.name + " " + p.lastName
}

func (s *student) changeName(name string) {
	s.name = name
}

type class struct {
	id       int
	name     string
	students []student
}

func toShower(x shower) {
	fmt.Println(x.toString())
}

func main() {

	student01 := student{}
	student01.id = 1
	student01.name = "Fulano"
	student01.lastName = "da Silva"
	student01.note = "B"

	student02 := student{}
	student02.id = 2
	student02.name = "Beltrano"
	student02.lastName = "Pereira"
	student02.note = "B"

	classOne := class{
		id:       1,
		name:     "Primeira Turma",
		students: []student{student01, student02}}

	for _, student := range classOne.students {
		fmt.Println(student)
	}

	newStudent := student{}
	newStudent.id = 3
	newStudent.name = "Chaves"
	newStudent.lastName = "do Oito"
	newStudent.note = "B"

	fmt.Println(newStudent.getFullName())

	newStudent.name = "Kiko"
	// fmt.Println(newStudent.toString())

	newTeacher := teacher{}
	newTeacher.id = 4
	newTeacher.name = "Cotinha"
	newTeacher.lastName = "Teacher"

	// fmt.Println(newTeacher.toString())

	toShower(newStudent)
	toShower(newTeacher)
}
