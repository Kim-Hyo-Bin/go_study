package main

import "fmt"

// struct 정의
type animal struct {
    name string
    age  int
}

func main() {
    var p1 animal 
    p1 = animal{"cat", 7}
    fmt.Println(p1)

    p2 := animal{name: "bird", age: 1}
    fmt.Println(p2)
    
    p3 := animal{}
    p3.name = "Dog"
    p3.age = 5
    fmt.Println(p3)
    //init var
    p3 = animal{}
    p3.name = "dog" 
    fmt.Println(p3)
}
