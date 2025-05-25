package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	firstArray := readInput()
	secondArray := readInput()
	slices.Sort(secondArray)
	res := intersection(firstArray, secondArray)	

	if len(res) == 0 {
		fmt.Println("Empty intersection")
	} else {
		fmt.Println(res)
	}
}

func readInput() []int {
	stdin := bufio.NewReader(os.Stdin)
	input, _ := stdin.ReadString('\n')
	numbers := strings.Fields(input)

	res := convertToInt(numbers)

	return res
}

func convertToInt(s []string) []int {
	var res []int
	for _, val := range s {

		temp, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Invalid input")
			os.Exit(1)
		}
		res = append(res, int(temp))
	}
	return res	
}

func intersection(f, s []int) []int {
	set := make(map[int]bool)
	for _, val := range s {
		set[val] = true
	}

	var res []int
	for _, val := range f {
		if set[val] {
			res = append(res, val)
		}
	}
	return res
}