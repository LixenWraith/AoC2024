// AoC 2024, Day 9, Part 2, Lixen Wraith
package main

import (
	"fmt"
	"os"
)

// data types
type disk = []int

type diskMap = map[int]cluster
type cluster struct {
	fileId     int
	fileBlocks byte
	freeBlocks byte
}

const freeSpace = -1 // file id for free space

// disk checksum
func checksum(data disk) uint64 {
	var c uint64 = 0

	for i := 0; i < len(data); i++ {
		d := data[i]
		if d != freeSpace {
			c += uint64(i * data[i])
		}
	}

	return c
}

func compact(dm diskMap) diskMap {
	numClusters := len(dm)
	cdm := make(diskMap, numClusters)
	compactFileOrder := make([]int, numClusters)
	var targetFreeBlocks, sourceFileBlocks byte

	// TODO

	return diskMap
}

// convert diskmap into disk data blocks
func write(dm diskMap) disk {
	d := disk{}

	for i := 1; i <= len(dm); i++ {
		for j := 0; j < int(dm[i].fileBlocks); j++ {
			d = append(d, dm[i].fileId)
		}
		for j := 0; j < int(dm[i].freeBlocks); j++ {
			d = append(d, freeSpace)
		}
	}

	return d
}

// create a byte slice with values equal to numerical text input ('0' -> 0)
func parse(input []byte) diskMap {
	const clusterRecordSize = 2
	numClusters := len(input) / clusterRecordSize
	dm := make(diskMap, numClusters)

	for i := 0; i < numClusters; i++ {
		dm[i] = cluster{
			fileId:     i,
			fileBlocks: input[i] - byte('0'),
			freeBlocks: input[i+clusterRecordSize/2] - byte('0'),
		}
	}

	return dm
}

// read input file into a slice of byte
func inputFromFile(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func main() {
	// filename := "./input.txt"
	filename := "./example1.txt"
	// filename := "./example2.txt"
	// filename := "./example3.txt"
	input, err := inputFromFile(filename)
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}

	// a := parse(input)
	// b := unpack(a)
	// c := optimize(b)
	// d := checksum(c)
	// output := checksum(pack(optimize(unpack(parse(input))))

	// fmt.Println(output)

	fmt.Printf("Part 1 : compact checksum = %d\n", checksum(write(compact(parse(input)))))
}