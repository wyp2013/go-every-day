package main

import "fmt"

func spiralPrint(i, j, m, n int, mat[][]int, res*[]int) {
	// fmt.Println(i,j,m,n)

	// 打印第一行数据
	r := i
	c := j
	for ; c <= n; c++ {
		*res = append(*res, mat[r][c])
	}

	// 打印最后一列数据
	r++
	c = n
	for ; r <= m; r++ {
		*res = append(*res, mat[r][c])
	}

	// 打印最后一行数据，要注意行数大于1，否则重复打印
	r = m
	c = n-1
	if m - i > 0 {
		for ; c >= j; c-- {
			*res = append(*res, mat[r][c])
		}
	}

	// 打印第一列数据，要注意列数大于1，否则会重复打印
	r = m-1
	c = j
	if n - j > 0 {
		for ; r > i; r-- {
			*res = append(*res, mat[r][c])
		}
	}
}


func spiralTravel(i, j, m, n int, mat[][]int, res*[]int) {
	if i > m || j > n {
		return
	}

	spiralPrint(i, j, m, n, mat, res)

	i = i+1
	j = j+1
	m = m-1
	n = n-1
	spiralTravel(i, j, m, n, mat, res)
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0, m*n)

	spiralTravel(0, 0, m-1, n-1, matrix, &res)
	return res
}


func main() {
	matrix := [][]int{{1,2,3,5},{4,5,6,7},{7,8,9,10}}

	//matrix := [][]int{{1},{4},{7}}
	res := spiralOrder(matrix)

	fmt.Println(res)
}

