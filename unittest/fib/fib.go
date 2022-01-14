package fib

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * Fib(n-1)
}
