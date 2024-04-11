package main

import (
	"strconv"
)

func Countdown(n int) string {
	str := ""
	for i := n; i > 0; i -= 2 {
		str += strconv.Itoa(i) + ", "
	}
	newstr := ""
	newstr = str + "0!"
	return string(newstr)
}
