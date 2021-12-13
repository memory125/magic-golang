package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution1 = make(map[string]int, len(users))
	distribution2 = make(map[string]int, len(users))
)

func main() {
	left1, left2 := dispatchCoin()
	fmt.Println("剩下left1, left2：", left1, left2)
}

func dispatchCoin() (int, int) {
	left1 := coins
	left2 := coins
	if left1 <= 0 {
		fmt.Println("金币不足！！！！")
		return 0, 0
	}

	if left2 <= 0 {
		fmt.Println("金币不足！！！！")
		return 0, 0
	}

	// 使用for...range实现
	for _, name := range users {
		for _, ch := range name {
			switch ch {
			case 'e', 'E':
				distribution1[name] += 1
				left1 -= 1

			case 'i', 'I':
				distribution1[name] += 2
				left1 -= 2

			case 'o', 'O':
				distribution1[name] += 3
				left1 -= 3

			case 'u', 'U':
				distribution1[name] += 4
				left1 -= 4
			}

			if left1 <= 0 {
				fmt.Println("金币不足！！！！")
				return 0, 0
			}
		}
	}

	// 使用for...i实现
	// 1. 依次拿到姓名
	for i := 0; i < len(users); i++ {
		// 2. 拿到一个人的名字，根据分金币的规则分金币
		// 2.1 每个人分到的金币应该保存到distribution中
		user := make([]byte, 0, len(users[i]))
		user = append(user, users[i]...)
		fmt.Printf("user = %s\n", user)
		distribution2[users[i]] = 0
		fmt.Println("byte(e), byte(E)", byte('e'), byte('E'))
		fmt.Println("byte(i), byte(I)", byte('i'), byte('I'))
		fmt.Println("byte(o), byte(O)", byte('o'), byte('O'))
		fmt.Println("byte(u), byte(U)", byte('u'), byte('U'))
		for j := 0; j < len(user); j++ {
			fmt.Println("user[j]", user[j])
			/*
				第一种方式：if...else
			 */
			//if user[j] == byte('e') || user[j] == byte('E') {
			//	distribution2[users[i]] += 1
			//	left -= 1
			//} else if user[j] == byte('i') || user[j] == byte('I') {
			//	distribution2[users[i]] += 2
			//	left -= 2
			//} else if user[j] == byte('o') || user[j] == byte('O') {
			//	distribution2[users[i]] += 3
			//	left -= 3
			//} else if user[j] == byte('u') || user[j] == byte('U') {
			//	distribution2[users[i]] += 4
			//	left -= 4
			//}
			//
			//if left <= 0 {
			//	fmt.Println("金币不足！！！！")
			//	return 0
			//}

			/*
				第二种方式：switch...case
			 */
			switch user[j] {
			case byte('e'), byte('E'):
				distribution2[users[i]] += 1
				// 2.2 还要记录剩下的金币数
				left2 -= 1
				break

			case byte('i'), byte('I'):
				distribution2[users[i]] += 2
				left2 -= 2
				break

			case byte('o'), byte('O'):
				distribution2[users[i]] += 3
				left2 -= 3
				break

			case byte('u'), byte('U'):
				distribution2[users[i]] += 4
				left2 -= 4
				break

			default:
				fmt.Println(user[j])
				break
			}

			if left2 <= 0 {
				fmt.Println("金币不足！！！！")
				return 0, 0
			}
		}
	}


	// 3. 第2步完成之后就能得到最终每个人分的金币数和剩余金币数

	fmt.Println("distribution1----", distribution1)
	fmt.Println("distribution2----", distribution2)
	return left1, left2
}
