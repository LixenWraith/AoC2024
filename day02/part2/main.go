// AoC 2024, Day 2, Part 2, Lixen Wraith
package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func safeReport(report []int) bool {
	if len(report) == 0 {
		return false
	}
	rAsc := append([]int{}, report...)
	rDesc := append([]int{}, report...)

	sort.Ints(rAsc)
	sort.Sort(sort.Reverse(sort.IntSlice(rDesc)))

	isSafe := true
	// first check: skip the reports that don't have only decreasing or increasing numbers
	if !reflect.DeepEqual(rAsc, report) &&
		!reflect.DeepEqual(rDesc, report) {
		isSafe = false
	} else {
		// second check: skip the reports level distance to adjacent level is >3 or <1
		for i := 0; i < len(report); i++ {
			if i > 0 {
				d := math.Abs(float64(report[i] - report[i-1]))
				if d < 1 || d > 3 {
					isSafe = false
					break
				}
			}

			if i < len(report)-1 {
				d := math.Abs(float64(report[i] - report[i+1]))
				if d < 1 || d > 3 {
					isSafe = false
					break
				}
			}
		}
	}
	return isSafe
}

// safe counts the number of safe reports
func safe(reports [][]int) (int, error) {
	if len(reports) == 0 {
		return 0, fmt.Errorf("no reports found")
	}

	safeReports := 0
	for _, r := range reports {
		if safeReport(r) {
			safeReports++
			continue
		}

		for i := range r {
			subReport := append(append([]int{}, r[:i]...), r[i+1:]...)
			if safeReport(subReport) {
				safeReports++
				break
			}
		}
	}
	return safeReports, nil
}

// parse converts the strictly structured test file into 2 int slices
func parse(filename string) ([][]int, error) {
	// read file into []byte
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	reportsRaw := strings.Split(string(content), "\n")

	reports := make([][]int, len(reportsRaw))
	for i := 0; i < len(reportsRaw); i++ {
		sr := strings.Split(strings.TrimSpace(reportsRaw[i]), " ")
		for l := range sr {
			ll, err := strconv.Atoi(sr[l])
			if err != nil {
				return nil, err
			}
			reports[i] = append(reports[i], ll)
		}
	}

	return reports, nil
}

func main() {
	reports, err := parse("./input.txt")
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}

	safeReports, err := safe(reports)
	if err != nil {
		fmt.Printf("Error processing reports: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part 2 : Safe reports = %v\n", safeReports)
}
