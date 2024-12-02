// AoC 2024, Day 2, Part 1, Lixen Wraith
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

// safe counts the number of safe reports
func safe(reports [][]int) (int, error) {
	if len(reports) == 0 {
		return 0, fmt.Errorf("no reports found")
	}
	safeReports := 0
	for r := range reports {
		if len(reports[r]) == 0 {
			return safeReports, fmt.Errorf("empty report")
		}
		rAsc := append([]int{}, reports[r]...)
		rDesc := append([]int{}, reports[r]...)

		sort.Ints(rAsc)
		sort.Sort(sort.Reverse(sort.IntSlice(rDesc)))

		// first check: skip the reports that don't have only decreasing or increasing numbers
		if !reflect.DeepEqual(rAsc, reports[r]) &&
			!reflect.DeepEqual(rDesc, reports[r]) {
			continue
		} else {
			unsafe := false
			// second check: skip the reports level distance to adjacent level is >3 or <1
			for i := 0; i < len(reports[r]); i++ {
				if i > 0 {
					d := math.Abs(float64(reports[r][i] - reports[r][i-1]))
					if d < 1 || d > 3 {
						unsafe = true
						break
					}
				}

				if i < len(reports[r])-1 {
					d := math.Abs(float64(reports[r][i] - reports[r][i+1]))
					if d < 1 || d > 3 {
						unsafe = true
						break
					}
				}
			}
			if !unsafe {
				safeReports++
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

	fmt.Printf("Part 1 : Safe reports = %v\n", safeReports)
}
