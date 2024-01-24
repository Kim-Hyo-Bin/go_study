package main

import (
	"fmt"
	"errors"
)

func safeDivide(dividend, divisor int) (int, error) {
    if divisor == 0 {
        return 0, errors.New("division by zero")
    }
    return dividend / divisor, nil
}

func main(){
result, err := safeDivide(8, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
}
