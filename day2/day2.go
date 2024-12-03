package day2

import (
	"advent-of-code-2024/day1"
	"bufio"
	"log"
	"math"
	"strconv"
	"strings"
)

func ConvertInput(input string) [][]int {
	result := make([][]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		var tmp_slice []int
		for _, number := range numbers {
			tmp, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			tmp_slice = append(tmp_slice, tmp)
		}
		result = append(result, tmp_slice)
	}
	return result
}

func AscendingOrDescending(input [][]int) []int {
	var result []int
	for _, slice := range input {
		ascending := true
		descending := true
		limit_exceeded := false
		for i := 1; i < len(slice); i++ {
			if slice[i] > slice[i-1] {
				descending = false
			} else if slice[i] < slice[i-1] {
				ascending = false
			} else if slice[i] == slice[i-1] {
				ascending = false
				descending = false
			}

			if math.Abs(float64(slice[i]-slice[i-1])) > 3 {
				limit_exceeded = true
			}
		}

		if ascending && !limit_exceeded {
			result = append(result, 1)
		} else if descending && !limit_exceeded {
			result = append(result, 1)
		} else {
			result = append(result, 0)
		}
	}
	return result
}

func Day2Part1(filename string) int {
	input := day1.ReadFile(filename)
	converted_input := ConvertInput(input)
	result_slice := AscendingOrDescending(converted_input)
	return day1.SliceSum(result_slice)
}
