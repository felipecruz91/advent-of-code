package main

import (
	"bufio"
	"io"
	"log"
	"slices"
	"strconv"
	"strings"
)

// safeReports returns the number reports which are safe.
func safeReports(r io.Reader) int {
	count := 0
	lines := readerToStrings(r)

	for _, line := range lines {
		levels := getLevels(line)

		if isSafe(levels) {
			count++
			continue
		}
	}

	return count
}

func dampenerSafeReports(r io.Reader) int {
	count := 0
	lines := readerToStrings(r)

out:
	for _, line := range lines {
		levels := getLevels(line)

		if isSafe(levels) {
			count++
			continue
		}

		for i := 0; i < len(levels); i++ {
			e := removeLevels(levels, []int{i})
			if isSafe(e) {
				count++
				continue out
			}
		}
	}

	return count
}

// removeLevels deletes from the levels slice the elements presents that have the given indices.
func removeLevels(levels []int, indices []int) []int {
	if len(levels) == 0 {
		return nil
	}

	e := make([]int, len(levels))
	copy(e, levels)

	set := make(map[int]bool)
	for _, v := range indices {
		set[v] = true
	}

	i := 0
	return slices.DeleteFunc(e, func(t int) bool {
		defer func() {
			i++
		}()
		return set[i]
	})
}

func isSafe(levels []int) bool {
	increasing := levels[1] > levels[0] // initial tendency is increasing

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		// checks that levels are either gradually increasing or gradually decreasing.
		if increasing && diff <= 0 {
			return false
		}
		if !increasing && diff >= 0 {
			return false
		}
		// check that any two adjacent levels differ by at least one and at most three.
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func readerToStrings(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

// getLevels returns the list of levels from a report line.
func getLevels(line string) []int {
	s := strings.Split(line, " ")

	levels := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		level, err := strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}
		levels = append(levels, level)
	}

	return levels
}
