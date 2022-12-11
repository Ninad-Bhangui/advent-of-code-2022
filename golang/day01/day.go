package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Ninad-Bhangui/advent-of-code-2022/internal/filereader"
)

func Calc1(inputFilepath string) int {
	elvesNumbers := findTotalByElf(inputFilepath)
	return findMax(elvesNumbers)
}

func Calc2(inputFilepath string) int {
	elvesNumbers := findTotalByElf(inputFilepath)
	sort.Sort(sort.Reverse(sort.IntSlice(elvesNumbers)))
	sum := 0
	for _, current := range elvesNumbers[0:3] {
		sum += current
	}
	return sum
}

func findElfSum(elfData []string) int {
	sum := 0
	for _, current := range elfData {
		currentNum, _ := strconv.Atoi(current)
		sum += currentNum
	}
	return sum
}
func findTotalByElf(inputFilepath string) []int {
	data := filereader.ReadFileToString(inputFilepath)
	elvesData := strings.Split(data, "\n\n")
	var elvesNumbers []int
	for _, elf := range elvesData {
		elfNumbers := strings.Split(elf, "\n")
		elfSum := findElfSum(elfNumbers)
		elvesNumbers = append(elvesNumbers, elfSum)
	}
	return elvesNumbers
}
func findMax(elfNumbers []int) int {
	max := 0
	for _, current := range elfNumbers {
		if current > max {
			max = current
		}
	}
	return max
}

func Main() {
	fmt.Println("day1: ", Calc1("inputs/input.txt"))
}
