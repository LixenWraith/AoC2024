// AoC 2024, Day 9, Part 1, Lixen Wraith
package main

import (
	"fmt"
	"os"
)

// data types
type data = []int
type index = []byte

const (
	recordLength = 2  // 1 digit for file blocks, 1 digit for free blocks
	freeSpace    = -1 // file id for free space
)

// disk checksum
func checksum(disk data) uint64 {
	var c uint64 = 0
	for i := 0; i < len(disk); i++ {
		d := disk[i]
		if d != freeSpace {
			c += uint64(i * disk[i])
		}
	}

	return c
}

// move all file blockss to the beginning of the disk
func compact(disk data) data {
	j := len(disk)
	d := make(data, j)
	for i := 0; i < j; i++ {
		d[i] = freeSpace
		block := disk[i]

		for block == freeSpace && i < j {
			j--
			block = disk[j]
		}

		d[i] = block
	}
	for j < len(disk) {
		d[j] = freeSpace
		j++
	}

	return d
}

// convert diskmap into disk data blocks
func unpack(diskMap index) (disk data) {
	d := data{}
	var fileRecordBlock, freeRecordBlock byte
	fileId := freeSpace

	for i := 0; i < len(diskMap)/recordLength; i++ {
		fileId++

		fileRecordBlock = diskMap[i*2]
		writeDisk(&d, fileId, fileRecordBlock)

		freeRecordBlock = diskMap[i*2+1]
		writeDisk(&d, freeSpace, freeRecordBlock)
	}

	if len(diskMap)%recordLength != 0 {
		writeDisk(&d, fileId+1, diskMap[len(diskMap)-1])
	}

	return d
}

func writeDisk(disk *data, id int, blocks byte) {
	for diskWriter := 0; diskWriter < int(blocks); diskWriter++ {
		*disk = append(*disk, id)
	}
}

// create a byte slice with values equal to numerical text input ('0' -> 0)
func parse(input []byte) index {
	l := make(index, len(input))

	for i, c := range input {
		l[i] = c - byte('0')
	}

	return l
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
	filename := "./input.txt"
	// filename := "./example1.txt"
	// filename := "./example2.txt"
	// filename := "./example3.txt"
	input, err := inputFromFile(filename)
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 : compact checksum = %d\n", checksum(compact(unpack(parse(input)))))
}