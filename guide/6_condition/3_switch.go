package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
 * 除非以 fallthrough 语句结束，否则分支会自动终止
 *
 * 没有条件的 switch 同 `switch true` 一样。
 * 这一构造使得可以用更清晰的形式来编写长的 if-then-else 链
 */
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
	switchOrder()
}

func switchOrder() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
