package selftest

import (
	"fmt"
	"testing"
)

func TestShufflingMachine(t *testing.T) {
	n := 2
	arr := []int{36, 52, 37, 38, 3, 39, 40, 53, 54, 41, 11, 12, 13, 42, 43, 44, 2, 4, 23, 24, 25, 26, 27, 6, 7, 8, 48, 49, 50, 51, 9, 10, 14, 15, 16, 5, 17, 18, 19, 1, 20, 21, 22, 28, 29, 30, 31, 32, 33, 34, 35, 45, 46, 47}

	card := shufflingMachine(n, arr)
	for i := range card {
		fmt.Printf("%s ", card[i])
	}
}
