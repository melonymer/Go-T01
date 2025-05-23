package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	x := getOperand("Input left operand")
	o := getOperation()
	y := getOperand("Input right operand")	

	res, err := calculating(x, y, o)

	switch err {
	case 3:
		fmt.Println("division by zero")
	case 0:
		fmt.Printf("Result = %.3f\n", res)
	}
}

func getOperand(s string) float64 {
	for {
		var input string
		fmt.Println(s)
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		return value
	}
}

func getOperation() string {
	for {
		var input string
		fmt.Print("Enter operation (+, -, *, /): ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		input = strings.TrimSpace(input)
		if input == "+" || input == "-" || input == "*" || input == "/" {
			return input
		}

		fmt.Println("Invalid input")
	}
}


func calculating(x, y float64, o string) (float64, int) {
	var res float64
	err := 0
	switch o {
	case "+" :
		res = x + y
	case "-" :
		res = x - y
	case "/" :
		if y == 0 {
			err = 3
		}
		res = x / y
	case "*" :
		res = x * y
	}	
	
	return res, err
}