package main

import "fmt"

type Life struct {
    Width, Height int
}

func main() {
    // 변수 life를 선언하고 Life 구조체의 포인터로 초기화
    life := &Life{Width: 10, Height: 5}

    // 포인터를 통해 구조체의 필드에 접근
    fmt.Printf("Width: %d, Height: %d\n", life.Width, life.Height)

    fmt.Println("\n------\n")

    // 변수 x를 선언하고 값을 할당
    x := 42
    fmt.Println("\nx:=42\n")

    // 변수 x의 메모리 주소를 출력
    fmt.Println("just x (x) :", x)
    fmt.Println("Address of x (&x):", &x)
    fmt.Println("value of address x (*&x) :", *&x)

    // 포인터 p를 선언하고 변수 x의 메모리 주소를 할당
    p := &x
    fmt.Println("\np := &x\n")

    // 포인터를 통해 변수 x의 값을 참조
    fmt.Println("just p (p) :", p)
    fmt.Println("Value at address p (*p) :", *p)
    fmt.Println("address p (&p) :", &p)
    fmt.Println("address p of address x (&*p) :", &*p)


    
}
