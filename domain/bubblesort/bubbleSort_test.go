package bubblesorter

import "testing"

func TestBubbleSortInts(t *testing.T) {
	input := []int{4, 1, 3, 2}
	out := BubbleSortInts(input)

	ans := []int{1, 2, 3, 4}
	for i := range input {
		if out[i] != ans[i] {
			t.Errorf("incorrect sort expect:%#v actual:%#v", ans, out)
			break
		}
	}
}

func TestBubbleSortText(t *testing.T) {
	input := []string{"444", "222", "111", "333"}
	out := BubbleSortTexts(input)

	ans := []string{"111", "222", "333", "444"}
	for i := range input {
		if out[i] != ans[i] {
			t.Errorf("incorrect sort expect:%#v actual:%#v", ans, out)
			break
		}
	}
}
