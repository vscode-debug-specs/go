package usecase

import bubblesorter "github.com/vscode-debug-specs/go/domain/bubblesort"

func SortText(input []string) []string {
	return bubblesorter.BubbleSortTexts(input)
}
