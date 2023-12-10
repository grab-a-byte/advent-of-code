package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFileAsLines(file *os.File) []string {
	var lines []string
	r := bufio.NewReader(file)
	s, e := readln(r)
	for e == nil {
		lines = append(lines, s)
		s, e = readln(r)
	}

	return lines
}

func readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func StrToIntsSpace(s string) []int {
	values := strings.Split(s, " ")
	output := make([]int, len(values))

	for i, v := range values {
		//Intentially ignor error, should wqork panic fine if not true
		val, _ := strconv.Atoi(v)
		output[i] = val
	}
	return output
}
