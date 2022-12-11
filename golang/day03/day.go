package day01

import (
	"bufio"
	"fmt"
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
	var totalPriority int
	for scanner.Scan() {
		line := scanner.Text()
		part1 := line[0 : len(line)/2]
		part2 := line[len(line)/2:]
		commonElements := findCommonElements(part1, part2)
		totalPriority += findPriorityFromRune(commonElements[0])
	}
	return totalPriority
}

func Calc2(inputFilepath string) int {
	file, err := os.Open(inputFilepath)
	defer file.Close()
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)
	var totalPriority int
	var group []string
	var index int
	for scanner.Scan() {
		line := scanner.Text()
		group = append(group, line)
		if (index+1)%3 == 0 {
			commonElementsFirst := findCommonElements(group[0], group[1])
			commonElements := findCommonElements(group[2], string(commonElementsFirst))
			totalPriority += findPriorityFromRune(commonElements[0])
			group = nil
		}
		index++
	}
	return totalPriority
}

func findCommonElements(list1, list2 string) []rune {
	var runes []rune
	for _, e := range list1 {
		if strings.ContainsRune(list2, e) {
			runes = append(runes, e)
		}
	}
	return runes
}
func findPriorityFromRune(character rune) int {
	priority, _ := strconv.Atoi(fmt.Sprintf("%d", character))
	if priority >= 97 && priority <= 122 {
		return priority - 96
	} else {
		return priority - 64 + 26
	}
}
