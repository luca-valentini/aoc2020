package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	min      int
	max      int
	letter   string
	password string
}

type passwordPolicyNew struct {
	positionOne int
	positionTwo int
	letter      string
	password    string
}

func (p passwordPolicy) isValid() bool {
	count := strings.Count(p.password, p.letter)
	if count <= p.max && count >= p.min {
		return true
	}
	return false
}

func (p passwordPolicyNew) isValid() bool {
	var matchOne bool
	var matchTwo bool
	if len(p.password) >= p.positionOne {
		matchOne = p.letter == string(p.password[p.positionOne-1])
	} else {
		matchOne = false
	}
	if len(p.password) >= p.positionTwo {
		matchTwo = p.letter == string(p.password[p.positionTwo-1])
	} else {
		matchTwo = false
	}
	if matchOne != matchTwo {
		return true
	}
	return false
}

func rowToPasswordPolicy(row string) passwordPolicy {
	s := strings.Fields(row)
	minMax := strings.Split(s[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])
	letter := strings.Trim(s[1], ":")
	return passwordPolicy{min: min, max: max, letter: letter, password: s[2]}
}

func rowToPasswordPolicyNew(row string) passwordPolicyNew {
	s := strings.Fields(row)
	oneTwo := strings.Split(s[0], "-")
	one, _ := strconv.Atoi(oneTwo[0])
	two, _ := strconv.Atoi(oneTwo[1])
	letter := strings.Trim(s[1], ":")
	return passwordPolicyNew{positionOne: one, positionTwo: two, letter: letter, password: s[2]}
}

func loadPasswords(scanner *bufio.Scanner) int {
	counter := 0
	for scanner.Scan() {
		row := scanner.Text()
		passwordPolicy := rowToPasswordPolicy(row)
		if passwordPolicy.isValid() {
			counter++
		}
	}
	return counter
}

func loadPasswordsNew(scanner *bufio.Scanner) int {
	counter := 0
	for scanner.Scan() {
		row := scanner.Text()
		passwordPolicy := rowToPasswordPolicyNew(row)
		fmt.Println(passwordPolicy)
		if passwordPolicy.isValid() {
			counter++
		}
	}
	return counter
}

func main() {
	input, err := os.Open("./day2/input.txt")
	if err != nil {
		panic(err)
	}
	// scanner := bufio.NewScanner(input)
	// result := loadPasswords(scanner)
	scanner := bufio.NewScanner(input)
	result := loadPasswordsNew(scanner)
	fmt.Println(result)
}
