package main

import (
	"fmt"
)

type Farm interface {
	price() int
	alive() bool
}

type Animal struct {
	name string
	age  int
	meat int
}

type Fish struct {
	name string
}

func (ani Animal) price() int {
	return ani.meat * 1000
}
func (ani Animal) alive() bool {
	return false
}

func (fish Fish) price() int {
	return 500
}

func (fish Fish) alive() bool {
	return true
}

func showPrice(ani ...Farm) {
	for _, s := range ani {
		tmp := s.price()
		fmt.Println(s, tmp)

	}
}
func main() {
	cow1 := Animal{"1++", 3, 800}
	cow2 := Animal{"1+", 1, 600}
	fish1 := Fish{"fisssh"}
	fish2 := Fish{"fisssssssh"}

	showPrice(cow1, cow2, fish1, fish2)

}
