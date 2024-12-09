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

	DO   = "do()"
	DONT = "don't()"
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

	return getRealMultiInstructions(bytes.NewReader(removeDisabledInstructions(b)))
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

func removeDisabledInstructions(b []byte) []byte {
	res := make([]byte, 0)
	s := string(b)
	doCopy := true

	for i := 0; i < len(s); i++ {
		if i+len(DONT) > len(s) {
			res = append(res, s[i:]...)
			break
		}
		if s[i:i+len(DONT)] == DONT {
			i += len(DONT)
			doCopy = false
			continue
		}
		if s[i:i+len(DO)] == DO {
			i += len(DO)
			doCopy = true
		}
		if doCopy {
			res = append(res, s[i])
		}
	}
	return res
}
