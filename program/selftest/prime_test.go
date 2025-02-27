package selftest

import "testing"

func TestPrime(t *testing.T) {
	t.Log(prime(10))
}

func TestGetPrimeNums(t *testing.T) {
	t.Log(findPrimePairs(getPrimeNums(5)))
}

func TestIsPrime(t *testing.T) {
	t.Log(isPrime(2))
}
