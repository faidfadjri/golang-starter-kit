package main

import "fmt"

type Anime interface {
	setName(name string) 
	getName() string
	setGenre(genre string)
	getGenre() string
}

type anime struct {
	name  string
	genre string
}

func (a *anime) setName(name string) {
	a.name = name
}

func (a *anime) getName() string {
	return a.name
}

func (a *anime) setGenre(genre string) {
	a.genre = genre
}

func (a *anime) getGenre() string {
	return a.genre
}

func main() {
	fmt.Println("Testing")
	a := &anime{}

	a.setName("Naruto")
	a.setGenre("Shounen")

	fmt.Printf("%s - %s", a.getName(), a.getGenre())
}
