package main

import (
	"container/list"
	"strconv"
)

// https://leetcode-cn.com/problems/basic-calculator/



func calculate(s string) int {
	chStk := list.New()
	numStk := list.New()

	for k := 0; k < len(s) ; {
		// 如果是数字
		if s[k] >= '0' && s[k] <= '9' {
			i := k
			for ; s[i] >= '0' && s[i] <= '9' && i < len(s) ; i++ {}
			str := s[k:i]
			num, _:= strconv.Atoi(str)

			if chStk.Len() == 0 || numStk.Len() == 0 {
				numStk.PushBack(num)
			} else {
				cel := chStk.Front()
				ch := cel.Value.(byte)

				if ch == '-' {

				} else if ch == '+' {

				} else {

				}
			}
			k = i
		} else if s[k] == '-' {

		} else if s[k] == '+' {

		} else if s[k] == '(' {

		} else if s[k] == '(' {

		} else {
			k++
		}
	}

	return -1
}

func main() {

}
