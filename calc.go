package main

import "fmt" // базовый вывод
import "os" // для получения параметров командной строки
import "strconv" // конвертация типов

func main() {
	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		panic(err)
	}
	
	action := os.Args[2]

	if action == "+" {
		var result float64 = a + b
		fmt.Println( result )
	}
	if action == "-" {
		var result float64 = a - b
		fmt.Println( result )
	}
	if action == "/"  {
		var result float64 = a / b
		fmt.Println( result )
	}
	if action == "*"  {
		var result  = a * b
		fmt.Println( result )
	}
}
