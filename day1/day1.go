package day1

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func SlicesDiff(a, b []int) []int {
	var diff []int
	for i, number := range a {
		if number == b[i] {
			diff = append(diff, 0)
		} else if number > b[i] {
			diff = append(diff, number-b[i])
		} else {
			diff = append(diff, b[i]-number)
		}
	}
	return diff
}

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func ConvertInput(input string) ([]int, []int) {
	var a, b []int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		tmpa, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		tmpb, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, tmpa)
		b = append(b, tmpb)
	}
	return a, b
}

func SliceSum(slice []int) int {
	sum := 0
	for _, number := range slice {
		sum += number
	}
	return sum
}

func CountNumber(slice []int, number int) int {
	count := 0
	for _, n := range slice {
		if n == number {
			count++
		}
	}
	return count
}

func Day1Part1(filename string) int {
	input := ReadFile(filename)
	a, b := ConvertInput(input)
	slices.Sort(a)
	slices.Sort(b)
	diff := SlicesDiff(a, b)
	return SliceSum(diff)
}

func Day1Part2(filename string) int {
	input := ReadFile(filename)
	a, b := ConvertInput(input)
	var nums []int
	for _, number := range a {
		nums = append(nums, CountNumber(b, number)*number)
	}
	return SliceSum(nums)
}
