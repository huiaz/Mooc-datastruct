package selftest

// 给定一个数字 n, 获取小于 n 的所有素数
func getPrimeNums(n int) []int {
	arrs := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			arrs = append(arrs, i)
		}
	}
	return arrs
}

// 判断一个数字是否为素数
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrimePairs(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	var count int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == 2 {
			count++
		}
	}

	return count
}

func prime(n int) int {
	arrs := getPrimeNums(n)
	return findPrimePairs(arrs)
}
