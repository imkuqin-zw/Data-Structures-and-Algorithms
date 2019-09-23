package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	pre1, pre2 := 0, 1
	for i := 2; i <= n; i++ {
		pre2 = pre1 + pre2
		pre1 = pre2 - pre1
	}
	return pre2
}
