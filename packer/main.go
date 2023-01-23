package main

import (
	"fmt"
	"packer/number"
	"packer/stringutil"
	"packer/stringutil/greeting"
)

func main() {
	fmt.Println(number.IsPrime(19))
	stringutil.Reserve()
	fmt.Println(greeting.WelcomeText)
}
