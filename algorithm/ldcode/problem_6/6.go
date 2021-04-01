package main

import "fmt"

/*

0     6      12      18      24      30
1   5 7   11 13   17 19   23 25   29 31
2 4   8 10   14 16   20 22   26 28   32
3     9      15      21      27      33

n=4,i是字符串第i索引
从上面的排列中可以看出:
    i%(2k)(n-1)==0   都在第0行放着，比如0,6(k=1),12(k=2),18(k=3),
    i%(2k+1)(n-1)==0 都在第最后行， 比如3,9(k=1),15(k=2),21(k=3),
    2k(n-1) =< i <= (2k+1)(n-1)   从上到下顺序放，比如0~3，6~9(k=1)，12~15(k=2)
    (2k+1)(n-1) < i < 2(k+1)(n-1）按照斜对角线从下到上放，比如4~5(k=0)，10~11(k=1)
安照这个思路，模拟法就可以搞出，模拟（从上到下，从下到上）
*/


func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	var sArr [][]byte
	for i := 0; i < numRows; i++ {
		sArr = append(sArr, make([]byte, 0))
	}

	l := len(s)
	r := 0    // 行数控制
	k := 0    // 上述的k，
	cnt := 0  // 控制k，cnt == 2(n-1)-2,k=k+1，表示新一轮开始

	for i := 0; i < l; i++ {
		// 从上到下，r++
		if i >= 2*k*(numRows-1) && i <= (2*k+1)*(numRows-1) {
			if i == 2*k*(numRows-1) {
				r = 0     // 从上到下，r=0开始
			}

			sArr[r] = append(sArr[r], s[i])
			r++
			cnt++

			if i == (2*k+1)*(numRows-1) {
				r = numRows-2 // 从下到上，r=numRows-2开始
			}
		} else if i > (2*k+1)*(numRows-1) && i < 2*(k+1)*(numRows-1) {
			// 从下到上，r--，注意r需要先减
			sArr[r] = append(sArr[r], s[i])
			r--
			cnt++
		}

		// 判断是否开始新一轮
		if cnt == 2*numRows-2 {
			k++
			cnt=0
		}
	}

	var buf []byte
	for i := 0; i < numRows; i++ {
		buf = append(buf, sArr[i]...)
	}

	return string(buf)
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3

	fmt.Println(convert(s,numRows))
}