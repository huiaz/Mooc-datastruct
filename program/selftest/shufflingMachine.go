package selftest

import "fmt"

func shufflingMachine(n int, arr []int) []string {
	// 洗牌机的逻辑
	// k 代表洗牌的次数，arr 代表洗牌的顺序
	card := []string{
		"S1", "S2", "S3", "S4", "S5", "S6", "S7", "S8", "S9", "S10", "S11", "S12", "S13",
		"H1", "H2", "H3", "H4", "H5", "H6", "H7", "H8", "H9", "H10", "H11", "H12", "H13",
		"C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "C10", "C11", "C12", "C13",
		"D1", "D2", "D3", "D4", "D5", "D6", "D7", "D8", "D9", "D10", "D11", "D12", "D13",
		"J1", "J2",
	}
	for k := 0; k < n; k++ {
		// 如果第i个位置的数字是j，则表示将第i张牌移到第j个位置
		// 洗牌的逻辑, 每次洗牌都需要新建一个数组，然后将原来的数组的值放到新数组的对应位置
		newCard := make([]string, len(card))

		for i := 0; i < len(arr); i++ {
			j := arr[i] // 获取第i个位置的数字
			// 将第i张牌移到第j个位置
			newCard[j-1] = card[i]
		}
		card = newCard

	}
	return card
}

func runSufflingMachine() {
	var n int
	var arr []int
	fmt.Scanf("%d", &n)
	for i := 0; i < 54; i++ {
		var num int
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}
	card := shufflingMachine(n, arr)
	for i := range card {
		// 打印数组，每个元素之间用空格隔开，最后一个元素后面没有空格
		if i == len(card)-1 {
			fmt.Printf("%s", card[i])
			break
		}
		fmt.Printf("%s ", card[i])
	}
}
