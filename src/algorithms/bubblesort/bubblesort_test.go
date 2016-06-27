package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	sortedValues := []int{1, 2, 3, 4, 5}
	BubbleSort(values)
	for index, val := range values {
		if val != sortedValues[index] {
			t.Error("BubbleSort() failed. Got ", values, "Expected 1 2 3 4 5")
      break;
		}
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	sortedValues := []int{1, 2, 3, 5, 5}
	BubbleSort(values)
	for index, val := range values {
		if val != sortedValues[index] {
			t.Error("BubbleSort() failed. Got ", values, "Expected 1 2 3 4 5")
      break;
		}
	}
}

func TestBubbleSort3(t *testing.T) {
	values := []int{5}
	sortedValues := []int{5}
	BubbleSort(values)
	for index, val := range values {
		if val != sortedValues[index] {
			t.Error("BubbleSort() failed. Got ", values, "Expected 1 2 3 4 5")
      break;
		}
	}
}
