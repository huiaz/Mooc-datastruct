package selftest

import (
	"fmt"
	"strings"
)

// 给定一个数字 n , 打印出一个 总数为<=n 的 沙漏形状, 例如:
// n = 17
// 打印出:
// *****
//
//	***
//	 *
//	***
//
// *****
// 注意: 每一行的 * 数量都是奇数, 并且每一行的 * 数量都是递增+2
func hourglassShape(N int, symbol string) {
	if N < 7 {
		fmt.Println(symbol)
		fmt.Println(N - 1)
		return
	}
	if N > 1000 {
		return
	}
	// 计算沙漏的一半高度的符号数
	halfN := N/2 + 1

	// 找出最大宽度
	maxWidth := 1
	halfNSymbols := 1
	for halfNSymbols+(maxWidth+2) <= halfN {
		maxWidth += 2
		halfNSymbols += maxWidth
	}

	// 如果当前总符号数多于N，减小最大宽度
	if halfNSymbols > halfN {
		maxWidth -= 2
		halfNSymbols -= maxWidth
	}

	// 打印沙漏的上半部分（包括中间行）
	for i := maxWidth; i >= 1; i -= 2 {
		// 计算当前行需要的空格数
		spaces := (maxWidth - i) / 2
		// 打印空格和符号
		fmt.Println(strings.Repeat(" ", spaces) + strings.Repeat(symbol, i))
	}

	// 打印沙漏的下半部分
	for i := 3; i <= maxWidth; i += 2 {
		// 计算当前行需要的空格数
		spaces := (maxWidth - i) / 2
		// 打印空格和符号
		fmt.Println(strings.Repeat(" ", spaces) + strings.Repeat(symbol, i))
	}

	// 如果沙漏的总符号数少于N，打印剩余的符号
	if halfNSymbols < halfN {
		fmt.Println(N - halfNSymbols*2 + 1)
	}
}
