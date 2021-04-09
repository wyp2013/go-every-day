package main

import (
	"fmt"
)

// 引用：https://mp.weixin.qq.com/s/MTZ0C9zYsNrb8wyIm2D8BA
func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	fmt.Println("---------init------------------")
	fmt.Println(slice)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))


	// s2、slice都发生改变
	s2 = append(s2, 100)
	fmt.Println("---------init2------------------")
	fmt.Println(slice)
	fmt.Println(s1)
	fmt.Println(s2, len(s2), cap(s2))
	//ptr :=(unsafe.Pointer)((uintptr)(unsafe.Pointer(&s1[0])) + unsafe.Offsetof(7))
    //testp := (*int)(ptr)
	//fmt.Println(*testp)

	fmt.Printf("地址相同 &slice[4]=%p, &s2[0]=%p \n", &slice[4], &s2[0])

	//  只有s2发生了改变，因为s2发生了扩容
	fmt.Println("---------init3------------------")
	s2 = append(s2, 200)
	fmt.Println(slice)
	fmt.Println(s1)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Printf("地址不同 &slice[4]=%p, &s2[0]=%p， s2发生了扩容 \n", &slice[4], &s2[0])


	// 只影响slice, s1, 不影响s2
	fmt.Println("---------init4------------------")
	s1[2] = 20
	fmt.Println(s1)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(slice)
}
