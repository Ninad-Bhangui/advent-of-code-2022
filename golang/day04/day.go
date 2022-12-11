package day01

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Calc1(inputFilepath string) int {
	file, err := os.Open(inputFilepath)
	defer file.Close()
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		part1 := strings.Split(line, ",")[0]
		part2 := strings.Split(line, ",")[1]
		part1Start, _ := strconv.Atoi(strings.Split(part1, "-")[0])
		part1End, _ := strconv.Atoi(strings.Split(part1, "-")[1])
		part2Start, _ := strconv.Atoi(strings.Split(part2, "-")[0])
		part2End, _ := strconv.Atoi(strings.Split(part2, "-")[1])

		if part1Start >= part2Start && part1End <= part2End {
			total += 1
		} else if part2Start >= part1Start && part2End <= part1End {
			total += 1
		}
	}
	return total
}

func Calc2(inputFilepath string) int {
	file, err := os.Open(inputFilepath)
	defer file.Close()
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		part1 := strings.Split(line, ",")[0]
		part2 := strings.Split(line, ",")[1]
		part1Start, _ := strconv.Atoi(strings.Split(part1, "-")[0])
		part1End, _ := strconv.Atoi(strings.Split(part1, "-")[1])
		part2Start, _ := strconv.Atoi(strings.Split(part2, "-")[0])
		part2End, _ := strconv.Atoi(strings.Split(part2, "-")[1])
		// 83-97,45-82
		// 2-4,6-8
		if part1Start >= part2Start && part1End <= part2End {
			total += 1
		} else if part2Start >= part1Start && part2End <= part1End {
			total += 1
		} else if (part1Start < part2Start) && (part1Start >= part2Start || part1End >= part2Start) {
			total += 1
		} else if (part2Start < part1Start) && (part2Start >= part1Start || part2End >= part1Start) {
			total += 1
		}
	}
	return total
}
