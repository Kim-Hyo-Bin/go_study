package main

import (
	"fmt"
)

type Animal struct {
	name string
	age, hand, foot int
}

func (ani Animal) sum_limbs() int {
	return ani.hand + ani.foot
}

func main() {
	dog := Animal{"happy", 5, 0, 4}

	fmt.Println(dog)
	fmt.Println(dog.sum_limbs())


	monkey := Animal{"marry", 3, 2, 2}

	fmt.Println(monkey)
	fmt.Println(monkey.sum_limbs())
}
