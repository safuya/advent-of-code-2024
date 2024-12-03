package day2

import (
	"bufio"
	"log"
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
