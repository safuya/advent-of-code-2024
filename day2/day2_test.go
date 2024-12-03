package day2

import (
	"reflect"
	"testing"
)

func TestConvertInput(t *testing.T) {
	input := `1 2 3 4
12 13 17 23 2 5
7 5 9`
	expected := [][]int{{1, 2, 3, 4}, {12, 13, 17, 23, 2, 5}, {7, 5, 9}}
	result := ConvertInput(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

//func TestAscendingOrDescending(t *testing.T) {
//input := day1.ReadFile("input_test.txt")
// converted_input := ConvertInput(input)
// result := AscendingOrDescending(input)
// expected := int[]{1, 0, 0, 0, 0, 1}

// if !reflect.DeepEqual(result, expected) {
//	t.Errorf("Expected %v, got %v", expected, result)
//}
//}
