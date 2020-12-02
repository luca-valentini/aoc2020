package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadExpenses(scanner *bufio.Scanner) ([]int, error) {
	expenses := []int{}
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}
func main() {
	input, err := os.Open("./day1/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)
	expenses, err := loadExpenses(scanner)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(expenses)-1; i++ {
		for j := i + 1; j < len(expenses); j++ {
			if expenses[i]+expenses[j] == 2020 {
				fmt.Println(expenses[i] * expenses[j])
			}
		}
	}

}
