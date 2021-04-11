package bubblesorter

import "strings"

// BubbleSort sorts the list
func BubbleSortInts(items []int) []int {
	length := len(items)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	return items
}

// BubbleSort sorts the list
func BubbleSortTexts(items []string) []string {
	length := len(items)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if strings.Compare(items[j], items[j+1]) > 0 {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	return items
}
