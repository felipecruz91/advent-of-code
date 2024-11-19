package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.OpenFile("input.txt", os.O_RDONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	v := run(f)

	fmt.Println(v)
}

func run(r io.Reader) int {
	sum := 0
	scanner := bufio.NewScanner(r)
	s := make([]rune, 2)

	for scanner.Scan() {
		row := scanner.Text()

		digits := make([]rune, 0, len(row))

		for i := 0; i < len(row); i++ {
			v := rune(row[i])
			if unicode.IsDigit(v) {
				digits = append(digits, v)
			}
		}

		s[0], s[1] = digits[0], digits[len(digits)-1]

		d, err := strconv.Atoi(string(s))
		if err != nil {
			log.Fatal(err)
		}
		sum += d
	}

	return sum
}
