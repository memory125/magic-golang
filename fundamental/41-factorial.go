package main

import "fmt"

// 递归
/*
   递归：函数自己调用自己
   递归适合处理那种相同问题的规模越来越小的场景
   递归一定要有一个明确的推出条件
 */

// 阶乘计算
/*
	3! = 3 * 2 * 1 = 3 * 2!
    4! = 4 * 3 * 2 * 1 = 4 * 3!
    5! = 5 * 4 * 3 * 2 * 1 = 5 * 4！
 */
func factorialCompute(n uint64) uint64  {
	if n <= 1 {
		return 1
	}

	return n * factorialCompute(n - 1)
}

// 走楼梯
/*
  n个台阶，一次可以走1步，也可以走2步，有多少种走法
 */
func stepsCompute(n uint64) uint64 {
	// 如果只有一个台阶
	if n == 1 {
		return 1
	}

	// 如果有两个台阶
	if n == 2 {
		return 2
	}

	return stepsCompute(n - 1) + stepsCompute(n - 2)
}


func main()  {
	/*
		7的阶乘！ 5040
		7个台阶的走法！ 21
	 */
	fmt.Println("7的阶乘！", factorialCompute(7))
	fmt.Println("7个台阶的走法！", stepsCompute(7))
}
