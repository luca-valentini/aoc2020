package main

import (
	"bufio"
	"errors"
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

func partOne(expenses []int) (int, error) {
	for i := 0; i < len(expenses)-1; i++ {
		for j := i + 1; j < len(expenses); j++ {
			if expenses[i]+expenses[j] == 2020 {
				return expenses[i] * expenses[j], nil
			}
		}
	}
	return 0, errors.New("Not found")
}

func partTwo(expenses []int) (int, error) {
	for i := 0; i < len(expenses)-2; i++ {
		for j := i + 1; j < len(expenses)-1; j++ {
			for k := j + 1; k < len(expenses); k++ {
				if expenses[i]+expenses[j]+expenses[k] == 2020 {
					return expenses[i] * expenses[j] * expenses[k], nil
				}
			}
		}
	}
	return 0, errors.New("Not found")
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
	res, err := partOne(expenses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1 result: %v\n", res)
	res, err = partTwo(expenses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2 result: %v\n", res)
}
