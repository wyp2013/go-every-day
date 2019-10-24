package sliceparam

import (
	"fmt"
	"testing"
)

// 值传递不会改变参数slice的长度，虽然在函数内部发生了扩容
func notModifySlice(s []int) {
	fmt.Println("-------------notModifySlice-------------------")
	fmt.Println("before ", len(s), cap(s))
	s = append(s, 4)
	fmt.Println("after ", len(s), cap(s))
	fmt.Println("-------------notModifySlice-------------------")
}

func TestNotMoidyfSlice(t *testing.T) {
	fmt.Println("------------------------TestNotModifySlice-------------------------")
	//s := []int{1, 2, 3}
	//
	//fmt.Println("before call notModifySlice ", len(s), cap(s))
	//notModifySlice(s)
	//fmt.Println("after call notModifySlice ", len(s), cap(s))

	ss := make([]int, 0, 10)
	ss = append(ss, 1)

	fmt.Println("before call notModifySlice ", ss, len(ss), cap(ss))
	notModifySlice(ss)

	//adx := unsafe.Pointer(&ss[0])
	//fmt.Println(int(adx)+ 1)
	fmt.Println("after call notModifySlice ", ss, len(ss), cap(ss))
}


// 指针传递会改变参数slice的长度
func modifySlice(s *[]int) {
	fmt.Println("-------------modifySlice-------------------")
	fmt.Println("before ", *s, len(*s), cap(*s))
	*s = append(*s, 4)
	fmt.Println("after ", *s, len(*s), cap(*s))
	fmt.Println("-------------modifySlice-------------------")
}

func TestMoidyfSlice(t *testing.T) {
	fmt.Println("------------------------TestModifySlice-------------------------")
	s := []int{1, 2, 3}

	fmt.Println("before call modifySlice ", s, len(s), cap(s))
	modifySlice(&s)
	fmt.Println("after call modifySlice ", s, len(s), cap(s))
}

// 虽然是值传递slice，但在函数内部修改了slice存放的元素值，会影响原来slice的元素，因为slice底层存放元素的成员是指针array(unsafepointer)
func modifySliceItem(s []int) {
	fmt.Println("-------------modifySliceItem------------------")
	s[0] = 100
	fmt.Println(s)
	fmt.Println("-------------modifySliceItem-------------------")
}

func TestMoidyfSliceItem(t *testing.T) {
	fmt.Println("------------------------TestModifySliceItem-------------------------")
	s := []int{1, 2, 3}

	fmt.Println("before call modifySliceItem ", s)
	modifySliceItem(s)
	fmt.Println("after call modifySlice ", s)
}


// 此列让你时刻记住slice是值专递，slice的len、cap
func TestAppend(t *testing.T) {
	a := []int{1,2,}
	fmt.Println("a", &a, len(a), cap(a))


	// a len 3, cap 4,
	a = append(a, 2)
	fmt.Println("a", &a, len(a), cap(a))


	var b, c []int
	// b len 4, cap 4; a len 3, cap 4, b.data = a.data
	b = append(a, 8)
	fmt.Println("a", &a, len(a), cap(a))
	fmt.Println("b", &b, len(b), cap(b))

	// c len 4, cap 4; a len 3, cap 4, c.data = a.data
	c = append(a, 32)
	fmt.Println("a", &a, len(a), cap(a))

	// so 最终b、c是一样
	fmt.Println(b)
	fmt.Println(c)
}





