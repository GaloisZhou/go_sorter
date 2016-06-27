package qsort

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for p <= j && temp <= values[j] {
			j--
		}
		if p <= j {
			values[p] = values[j]
			p = j
		}
		if i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}

	values[p] = temp

	if p-left > 1 {
		quickSort(values, left, p-1)
	}

	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
