package selftest

import "fmt"

func circleArr(arr []int, k, m int) {
	// k 代表数组的长度，m 代表向右移动的位数
	// 例如: arr = [1, 2, 3, 4, 5, 6, 7], k = 7, m = 3
	// 输出: [5, 6, 7, 1, 2, 3, 4]
	// 不使用额外的数组，直接在原数组上操作
	// 首先判断 k 和 m 是否合法
	if k < 1 || m < 1 {
		return
	}
	if m > k {
		m = m % k
	}

	for i := 0; i < m; i++ {
		// 循环移动数组最后几位到最前面, 使用 idx 记录当前需要移动的位置
		idx := k - m + i
		// 记录当前需要移动的位置的值
		tmp := arr[idx]
		// 从当前需要移动的位置开始，将前面的位置的值都往后移动一位
		for j := idx; j > i; j-- {
			arr[j] = arr[j-1]
		}
		// 将当前需要移动的位置的值放到最前面
		arr[i] = tmp
	}
}

func runCircleArr() {
	var k int
	var m int
	var arr []int
	fmt.Scanf("%d %d", &k, &m)
	for i := 0; i < k; i++ {
		var num int
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}
	for i := range k {
		// 打印数组，每个元素之间用空格隔开，最后一个元素后面没有空格
		if i == k-1 {
			fmt.Printf("%d", arr[i])
			break
		}
		fmt.Printf("%d ", arr[i])
	}
}
