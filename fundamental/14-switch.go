package main

import "fmt"

func main()  {
	statusSwtich(200)
	ageSwitch(28)
}

func statusSwtich(status int)  {
		switch status {
	case 200:
		fmt.Println("Ok!")
		break

	case 400:
		fmt.Println("Error!")
		break

	case 500:
		fmt.Println("Internal Error!")
		break

	default:
		fmt.Println("Nothing found!")
		break
	}
}

func ageSwitch(age int)  {
	switch {
	case age >= 60:
		fmt.Println("Older!")
		break

	case age < 60 && age >= 18:
		fmt.Println("Adult!")
		break

	case age < 18:
		fmt.Println("Yongerster!")
		break

	default:
		fmt.Println("Unknown!")
		break
	}
}
