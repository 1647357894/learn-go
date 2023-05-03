package main

import "fmt"

func main() {
	i := 6

	if i > 10 {

		fmt.Println("i>10")

	} else if i > 5 && i <= 10 {

		fmt.Println("5<i<=10")

	} else {

		fmt.Println("i<=5")

	}

	if i := 6; i > 10 {

		fmt.Println("i>10")

	} else if i > 5 && i <= 10 {

		fmt.Println("5<i<=10")

	} else {

		fmt.Println("i<=5")

	}

	switch i := 6; {

	case i > 10:

		fmt.Println("i>10")

	case i > 5 && i <= 10:

		fmt.Println("5<i<=10")

	default:

		fmt.Println("i<=5")

	}

	switch j := 1; j {

	case 1:
		// fallthrough 关键字 执行下一个紧跟着的case
		fallthrough

	case 2:

		fmt.Println("1")

	default:

		fmt.Println("没有匹配")

	}

	switch 2 > 1 {

	case true:

		fmt.Println("2>1")

	case false:

		fmt.Println("2<=1")

	}

	sum := 0

	for i := 1; i <= 100; i++ {

		sum += i

	}

	fmt.Println("the sum is", sum)

	sum = 0
	i = 1

	for i <= 100 {

		sum += i

		i++

	}

	fmt.Println("the sum is", sum)

	sum = 0
	i = 1

	//  continue、break
	for {

		sum += i

		i++

		if i > 100 {
			break
		}

	}

	fmt.Println("the sum is", sum)

}
