package main

import (
	"fmt"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "n sei"
	}
}

func main() {
	fmt.Println(tipo(2.3))
}
