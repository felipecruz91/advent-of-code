package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func distance(r io.Reader) (int, error) {
	l1, l2, err := getIntSlices(r)
	if err != nil {
		return -1, err
	}

	sort.Ints(l1)
	sort.Ints(l2)

	distance := 0
	for i := 0; i < len(l1); i++ {
		diff := l1[i] - l2[i]
		if diff < 0 {
			diff = -diff
		}
		distance += diff
	}
	return distance, nil
}

func score(r io.Reader) (int, error) {
	l1, l2, err := getIntSlices(r)
	if err != nil {
		return -1, err
	}

	m := make(map[int]int, len(l2)) // stores the frequency of each number
	for _, v := range l2 {
		m[v]++
	}

	score := 0
	for _, v := range l1 {
		score += v * m[v]
	}
	return score, nil
}

func getIntSlices(r io.Reader) ([]int, []int, error) {
	lines := readerToStrings(r)

	l1 := make([]int, 0, len(lines))
	l2 := make([]int, 0, len(lines))

	for _, line := range lines {
		s := strings.Split(line, "   ")
		if len(s) != 2 {
			return l1, l2, fmt.Errorf("couldn't split line by the given separator")
		}

		s0, err := strconv.Atoi(s[0])
		if err != nil {
			return l1, l2, err
		}
		l1 = append(l1, s0)

		s1, err := strconv.Atoi(s[1])
		if err != nil {
			return l1, l2, err
		}
		l2 = append(l2, s1)
	}

	return l1, l2, nil
}

func readerToStrings(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}
