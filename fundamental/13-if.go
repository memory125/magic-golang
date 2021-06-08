package main

import "fmt"

func main() {
	var age = 20

	if age > 18 {
		fmt.Println("Age is over 18!")
	} else if age > 9 {
		fmt.Println("Age is over 9!")
	} else {
		fmt.Println("Age is less than 9!!")
	}
}
