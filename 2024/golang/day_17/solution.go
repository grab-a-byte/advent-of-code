package day_17

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	content, err := os.ReadFile("./day_17/input.txt")
	if err != nil {
		panic("unable to open day 17 input")
	}

	strContent := string(content)
	answer1 := partOneSolution(strContent)
	fmt.Printf("Day 17 - Part 1: %s \n", answer1)
	answer2 := partTwoSolution(strContent)
	fmt.Printf("Day 17 - Part 2: %d \n", answer2)
}

func partOneSolution(input string) string {
	sanatixedInput := strings.ReplaceAll(input, "\r", "")
	var a, b, c int
	parts := strings.Split(sanatixedInput, "\n")
	fmt.Sscanf(parts[0], "Register A: %d", &a)
	fmt.Sscanf(parts[1], "Register B: %d", &b)
	fmt.Sscanf(parts[2], "Register C: %d", &c)

	programStr := strings.Split(strings.TrimLeft(parts[4], "Program: "), ",")
	program := make([]int, len(programStr))
	for i, s := range programStr {
		num, _ := strconv.Atoi(s)
		program[i] = num
	}

	machine := machine{
		a:       a,
		b:       b,
		c:       c,
		program: program,
		output:  []int{},
	}

	machine.run()
	var values []string
	for _, i := range machine.output {
		values = append(values, strconv.Itoa(i))
	}

	return strings.Join(values, ",")
}

func partTwoSolution(input string) int {
	sanatixedInput := strings.ReplaceAll(input, "\r", "")
	parts := strings.Split(sanatixedInput, "\n")[4]
	programStr := strings.Split(strings.TrimLeft(parts, "Program: "), ",")
	program := make([]int, len(programStr))
	for i, s := range programStr {
		num, _ := strconv.Atoi(s)
		program[i] = num
	}
	i := 0
	for {
		fmt.Println("processing" + string(i))
		m := machine{
			a:       i,
			program: program,
		}

		m.run()
		if testEq(m.output, m.program) {
			return i
		}
	}

	panic("No number found")
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type machine struct {
	a, b, c int
	program []int
	ip      int
	output  []int
}

func (m *machine) run() {
	for m.ip < len(m.program) {
		m.step()
	}
}

func (m *machine) step() {
	ins := m.currIns()
	m.ip++
	switch ins {
	case 0:
		operand := m.getOperand()
		div := int(math.Pow(2, float64(operand)))
		m.a = m.a / div
		m.ip++
	case 1:
		m.b = m.b ^ m.program[m.ip] //Literal operand
		m.ip++
	case 2:
		m.b = m.getOperand() % 8
		m.ip++
	case 3:
		if m.a == 0 {
			m.ip++
			break
		}
		m.ip = m.program[m.ip]
	case 4:
		m.b = m.b ^ m.c
		m.ip++
	case 5:
		val := m.getOperand() % 8
		m.output = append(m.output, val)
		m.ip++
	case 6:
		operand := m.getOperand()
		div := int(math.Pow(2, float64(operand)))
		m.b = m.a / div
		m.ip++
	case 7:
		operand := m.getOperand()
		div := int(math.Pow(2, float64(operand)))
		m.c = m.a / div
		m.ip++
	}
}

func (m *machine) currIns() int {
	return m.program[m.ip]
}

func (m *machine) getOperand() int {
	switch m.currIns() {
	case 0, 1, 2, 3:
		return m.currIns()
	case 4:
		return m.a
	case 5:
		return m.b
	case 6:
		return m.c
	}

	panic("Unknown operand number")
}
