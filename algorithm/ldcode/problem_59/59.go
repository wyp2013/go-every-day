package main

import "fmt"

func spiralPrint(i, j, m, n int, res*[][]int, cnt *int) {
	// fmt.Println(i,j,m,n)

	// 打印第一行数据
	r := i
	c := j
	for ; c <= n; c++ {
		*cnt++
		(*res)[r][c] = *cnt
	}

	// 打印最后一列数据
	r++
	c = n
	for ; r <= m; r++ {
		*cnt++
		(*res)[r][c] = *cnt
	}

	// 打印最后一行数据，要注意行数大于1，否则重复打印
	r = m
	c = n-1
	if m - i > 0 {
		for ; c >= j; c-- {
			*cnt++
			(*res)[r][c] = *cnt
		}
	}

	// 打印第一列数据，要注意列数大于1，否则会重复打印
	r = m-1
	c = j
	if n - j > 0 {
		for ; r > i; r-- {
			*cnt++
			(*res)[r][c] = *cnt
		}
	}
}


func spiralTravel(i, j, m, n int, res*[][]int, cnt *int) {
	if i > m || j > n {
		return
	}

	spiralPrint(i, j, m, n, res, cnt)

	i = i+1
	j = j+1
	m = m-1
	n = n-1
	spiralTravel(i, j, m, n, res, cnt)
}


func generateMatrix(n int) [][]int {
	var res [][]int
	for i := 0; i < n; i++ {
		line := make([]int, n)
		res  = append(res, line)
	}

	cnt := 0
	spiralTravel(0, 0, n-1, n-1, &res, &cnt)

	return res
}

func main() {
	res := generateMatrix(1)
	fmt.Print(res)

	res = generateMatrix(3)
	fmt.Print(res)
}
