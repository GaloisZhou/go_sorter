package qsort

import (
	"fmt"
	"testing"
)

func TestQuestSort1(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	sortedValues := []int{1, 2, 3, 4, 5, 6}
	QuickSort(values)

	fmt.Println(values)

	for i, val := range values {
		if val != sortedValues[i] {
			t.Error(values, sortedValues)
			break
		}
	}
}

func TestQuestSort2(t *testing.T) {
	values := []int{1, 8, 37, 41, 31, 5, 3, 6, 1, 8}
	sortedValues := []int{1, 1, 3, 5, 6, 8, 8, 31, 37, 41}
	QuickSort(values)

	fmt.Println(values)

	for i, val := range values {
		if val != sortedValues[i] {
			t.Error(values, sortedValues)
			break
		}
	}
}

func TestQuestSort3(t *testing.T) {
	values := []int{8, 7, 6, 5, 4, 3, 2, 1}
	sortedValues := []int{1, 2, 3, 4, 5, 6, 7, 8}
	QuickSort(values)

	fmt.Println(values)

	for i, val := range values {
		if val != sortedValues[i] {
			t.Error(values, sortedValues)
			break
		}
	}
}
