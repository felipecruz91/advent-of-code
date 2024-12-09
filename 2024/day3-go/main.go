package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"unicode"
)

const (
	OPERATION = "mul"
	LPAR      = "("
	COMMA     = ","
	RPAR      = ")"
)

func getRealMultiInstructions(r io.Reader) int {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	total := 0

	for {
		_, after, found := bytes.Cut(b, []byte(OPERATION+LPAR))
		if !found || len(after) == 0 {
			break
		}

		commaIdx := bytes.Index(after, []byte(COMMA))
		if commaIdx == -1 {
			fmt.Println("Comma not found, skipping to next")
			b = after
			continue
		}

		rParIdx := bytes.Index(after, []byte(RPAR))
		if rParIdx == -1 {
			fmt.Println("Right parenthesis not found, skipping to next")
			b = after
			continue
		}

		leftNum, err := getNumber(after, 0, commaIdx)
		if err != nil {
			fmt.Printf("Invalid left number: %s\n", err)
			b = after
			continue
		}

		rightNum, err := getNumber(after, commaIdx+1, rParIdx)
		if err != nil {
			fmt.Printf("Invalid right number: %s\n", err)
			b = after
			continue
		}

		total += leftNum * rightNum

		b = after
	}

	return total
}

func getRealMultiInstructionsPart2(r io.Reader) int {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	cpy := []byte{}
	s := string(b)
	doCopy := true
	for i := 0; i < len(s); i++ {
		if !doCopy && i+7 > len(s) {
			fmt.Println("End")
			break
		}
		if doCopy && i+4 > len(s) {
			fmt.Println("End")
			break
		}

		if i+7 > len(s) {
			cpy = append(cpy, s[i:]...)
			break
		}
		if s[i:i+7] == "don't()" {
			i += 7
			doCopy = false
			continue
		}
		if s[i:i+4] == "do()" {
			i += 4
			doCopy = true
		}

		if doCopy {
			cpy = append(cpy, s[i])
		}
	}

	b = cpy
	total := 0
	dont := false
	for {
		before, after, found := bytes.Cut(b, []byte(OPERATION+LPAR))
		if !found || len(after) == 0 {
			break
		}

		commaIdx := bytes.Index(after, []byte(COMMA))
		if commaIdx == -1 {
			fmt.Println("Comma not found, skipping to next")
			b = after
			continue
		}

		rParIdx := bytes.Index(after, []byte(RPAR))
		if rParIdx == -1 {
			fmt.Println("Right parenthesis not found, skipping to next")
			b = after
			continue
		}

		leftNum, err := getNumber(after, 0, commaIdx)
		if err != nil {
			b = after
			continue
		}

		rightNum, err := getNumber(after, commaIdx+1, rParIdx)
		if err != nil {
			b = after
			continue
		}
		if bytes.Contains(before, []byte("don't()")) {
			dont = true
			b = after
			continue
		}
		if bytes.Contains(before, []byte("do()")) {
			dont = false
		}
		if !dont {
			total += leftNum * rightNum
		}

		b = after
	}

	return total
}

// getNumber returns the number present within the two indexes.
func getNumber(b []byte, idx1, idx2 int) (int, error) {
	for i := idx1; i < idx2; i++ {
		if !unicode.IsDigit(rune(b[i])) {
			return -1, fmt.Errorf("%q is not a digit", string(string(b)[i]))
		}
	}

	n, err := strconv.Atoi(string(b)[idx1:idx2])
	if err != nil {
		return -1, fmt.Errorf("failed to convert %q to an int", string(b)[idx1:idx2])
	}
	return n, nil
}
