package main

// map面试题测试

const N = 3

func main() {
	m := make(map[int]*int)
	for i := 0; i < N; i++ {
		b := i
		m[i] = &b
	}
	for _, v := range m {
		print(*v)
	}
}
