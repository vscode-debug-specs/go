package usecase

import bubblesorter "github.com/vscode-debug-specs/go/domain/bubblesort"

func SortInts(input []int) []int {
	return bubblesorter.BubbleSortInts(input)
}
