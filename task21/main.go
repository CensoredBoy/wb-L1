package main

import "fmt"

type Movement interface {
	Move(distance int) string 
}

type Man struct {
	place int
}

func (m *Man) SpecificMove(dist int) string {
	m.place += dist
	return fmt.Sprintf("I walked %d kilometers", dist)
}

type ManAdapter struct {
	*Man
}

func (a *ManAdapter) Move(distance int) string {
	return a.SpecificMove(distance)
}

func main() {
	var m Movement
	m = &ManAdapter{&Man{12}}
	fmt.Println(m.Move(12))
}