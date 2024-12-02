package day1

import (
	"reflect"
	"testing"
)

func TestSlicesDiff(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 5, 7, 9, 11}
	expected := []int{2, 3, 4, 5, 6}

	diff := SlicesDiff(a, b)

	if !reflect.DeepEqual(diff, expected) {
		t.Errorf("Expected %v, got %v", expected, diff)
	}
}

func TestReadFile(t *testing.T) {
	result := ReadFile("input_test.txt")
	expected := `123   456
789   101112`

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestConvertInput(t *testing.T) {
	input := `123   456
789   101112`
	expecteda := []int{123, 789}
	expectedb := []int{456, 101112}

	a, b := ConvertInput(input)

	if !reflect.DeepEqual(a, expecteda) {
		t.Errorf("a: Expected %v, got %v", expecteda, a)
	}

	if !reflect.DeepEqual(b, expectedb) {
		t.Errorf("b: Expected %v, got %v", expectedb, b)
	}
}

func TestSliceSum(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	expected := 15

	sum := SliceSum(slice)

	if sum != expected {
		t.Errorf("Expected %d, got %d", expected, sum)
	}
}
