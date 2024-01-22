package main

import "fmt"

type MyType struct {
	Value int
}

func (v MyType) SetValue(newValue int) {
	v.Value = newValue
}

func (p *MyType) SetValueWithPointer(newValue int) {
	p.Value = newValue
}

func main() {
	myInstance := MyType{Value: 10}

	myInstance.SetValue(20)
	fmt.Println("Value after using Value Receiver:", myInstance.Value) // 출력: 10 (원본 구조체 변경되지 않음)

	myInstance.SetValueWithPointer(30)
	fmt.Println("Value after using Pointer Receiver:", myInstance.Value) // 출력: 30 (원본 구조체 변경됨)
}
