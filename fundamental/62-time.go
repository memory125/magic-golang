package main

import (
	"fmt"
	"time"
)

// 时间包

func main()  {
	now := time.Now()
	/*
		=========Time========
		2021-12-23 15:50:46.8068874 +0800 CST m=+0.007116101
		2021
		December
		23
		2021 December 23
		15
		50
		46
		1640245846
		1640245846806887400
	 */
	fmt.Println("=========Time========")
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	/*
		2021-12-23 15:50:46 +0800 CST
		2021
		December
		23
	*/
	timeSpan := time.Unix(1640245846, 0)
	fmt.Println(timeSpan)
	fmt.Println(timeSpan.Year())
	fmt.Println(timeSpan.Month())
	fmt.Println(timeSpan.Day())

	/*
		1s
		2021-12-24 16:05:17.8437475 +0800 CST m=+86400.008259401
	 */
	// 时间间隔
	fmt.Println(time.Second)

	// now + 24小时
	fmt.Println(now.Add(24 * time.Hour))

	// 定时器
	//timer := time.Tick(time.Second * 5)
	//for t := range timer {
	//	/*
	//		====== 2021-12-23 16:15:42.2176587 +0800 CST m=+5.041362701
	//		====== 2021-12-23 16:15:47.2240099 +0800 CST m=+10.047713901
	//		====== 2021-12-23 16:15:52.2249765 +0800 CST m=+15.048680501
	//		====== 2021-12-23 16:15:57.2259057 +0800 CST m=+20.049609701
	//		====== 2021-12-23 16:16:02.2270022 +0800 CST m=+25.050706201
	//		====== 2021-12-23 16:16:07.2246255 +0800 CST m=+30.048329501
	//		====== 2021-12-23 16:16:12.2187368 +0800 CST m=+35.042440801
	//	 */
	//	fmt.Println("======", t)   // 5秒钟执行一次
	//}

	// 格式化时间
	// 2021-12-23
	fmt.Println(now.Format("2006-01-02"))
	// 2021/12/23
	fmt.Println(now.Format("2006/01/02"))
	// 2021-12-23 16:18:49
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	// 2021-12-23 16:18:49 PM
	fmt.Println(now.Format("2006-01-02 15:04:05 PM"))
	// 2021-12-23 16:18:49.488
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))

	// 按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02 15:04:05", "2021-12-23 17:00:00")
	if err != nil {
		fmt.Printf("Parse time failed! error is %v.\n", err)
		return
	}
	fmt.Println("======Time Parse======")
	/*
		2021-12-23 17:00:00 +0000 UTC
		1640278800
	 */
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
}
