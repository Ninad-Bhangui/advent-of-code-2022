package day01

import (
	"fmt"
	"testing"
)

var sampleFilePath = "inputs/sample1.txt"
var inputFilePath = "inputs/input1.txt"

func assertIntEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d but got %d", want, got)
	}
}
func TestFirst(t *testing.T) {
	want := 15

	got := Calc1(sampleFilePath)
	assertIntEquals(t, got, want)
	fmt.Println("day01: ", Calc1(inputFilePath))
}

func TestSecond(t *testing.T) {
	want := 12
	got := Calc2(sampleFilePath)
	assertIntEquals(t, got, want)
	fmt.Println("day02: ", Calc2(inputFilePath))
}
