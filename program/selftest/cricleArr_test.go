package selftest

import (
	"fmt"
	"testing"
)

func TestCircleArr(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	k := 6
	m := 2
	circleArr(arr, k, m)
	for i := range k {
		// 打印数组，每个元素之间用空格隔开，最后一个元素后面没有空格
		if i == k-1 {
			fmt.Printf("%d", arr[i])
			break
		}
		fmt.Printf("%d ", arr[i])
	}

}

func TestCircleArrTest(t *testing.T) {
	runCircleArr()
}
