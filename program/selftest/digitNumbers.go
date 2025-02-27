package selftest

import (
	"fmt"
	"strings"
)

func digitNumbers(n string) map[int]int {
	// 使用map 结构存储每一个数字出现的次数
	digitMap := make(map[int]int)
	for i := 0; i < len(n); i++ {
		// 将数字转换成字符串
		digit := int(n[i])
		// 如果map 中已经存在该数字，将其出现的次数加1
		digitMap[digit]++
	}
	return digitMap
}

func compareDigitNumbers(n1, n2 string) bool {
	digitMap1 := digitNumbers(n1)
	digitMap2 := digitNumbers(n2)
	if len(digitMap1) != len(digitMap2) {
		return false
	}
	for k, v := range digitMap1 {
		if digitMap2[k] != v {
			return false
		}
	}
	return true
}

// 20位的正整数
func isValid(n string) bool {
	if len(n) > 20 {
		return false
	}
	for _, c := range n {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func doubleStringNumber(s string) string {
	if !isValid(s) {
		return ""
	}

	// 创建一个长度为len(s)的数组，初始化为0
	digits := make([]int, len(s))
	for i, c := range s {
		digits[i] = int(c - '0')
	}

	// 处理每一位数字 * 2，得到一个新的数字，使用[]int 存储
	carry := 0 // 进位标记
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i]*2 + carry
		if digits[i] > 9 {
			carry = 1
			digits[i] -= 10
		} else {
			carry = 0
		}
	}

	// 如果有进位，需要在前面加1
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}

	// 将[]int 转换成string
	var builder strings.Builder
	for _, digit := range digits {
		builder.WriteString(fmt.Sprintf("%d", digit))
	}

	return builder.String()
}

func runDigitNumbers() {
	var n string
	fmt.Scanf("%s", &n)

	n2 := doubleStringNumber(n)
	if !isValid(n) || len(n2) != len(n) {
		fmt.Println("No")
		fmt.Println(n2)
		return
	}
	if compareDigitNumbers(n, n2) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	fmt.Println(n2)
}
