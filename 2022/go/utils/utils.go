package utils

import (
	"bufio"
	"os"
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
