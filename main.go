package main

/*
File name    : main.go
Author       : miaoyc
Create date  : 2021/8/25 5:04 下午
Description  :
*/

import (
	"fmt"
	"github.com/miaoyc1989/python2go/python"
)


func main() {
	numbers := []int{0,1,2,3}
	fmt.Println(python.in(numbers, 1))
}