package main

import (
	"strconv"
)

/*
The ParseStringToInt function takes in a string (and some other parameters) and
returns a number num and an error err. You should inspect the err to
see if the ParseInt function returns anything. If there is an error,
you can handle it as you want. In this example, we panic, which exits
the program.

for example
{
str := "123456789"
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number is: ", num)
}
*/

func ParseStringToInt(str string) (num int64, err error) {
	num, err = strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return
}
