package main

import (
	"fmt"
	"reflect"
)

// sliceTest 学习Slice
func sliceTest() {
	fmt.Println("学习切片")
	arr := [...]int{1, 2, 3, 4, 5, 6, 8}
	s1 := arr[3:5] // 与Python相似
	s2 := arr[:5]
	s3 := arr[3:]
	s4 := arr[:]
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println("修改原数组前")
	fmt.Println("arr[3:5]", s1)
	fmt.Println("arr[:5]", s2)
	fmt.Println("arr[3:]", s3)
	fmt.Println("arr[:]", s4)
	fmt.Println("修改原数组后")
	updateSlice(arr[:])
	fmt.Println("arr[3:5]", s1)
	fmt.Println("arr[:5]", s2)
	fmt.Println("arr[3:]", s3)
	fmt.Println("arr[:]", s4)
	fmt.Println("可以发现，切片是引用类型，修改原数组，切片也会跟着变化")
}

// updateSlice 切片引用的证明
func updateSlice(s []int) {
	s[0] = 100
}

func initSlice() {
	var s1 []int                                                     // 声明一个空的切片
	fmt.Printf("s1 len:%d,cap%d,value:%v  \n", len(s1), cap(s1), s1) //[]
	for i := 0; i < 67; i++ {
		s1 = append(s1, i) // 追加元素,当容量不足时,会自动扩展长度(垃圾回收机制,生成新的变量,删除原变量,地址发生变化)
		if i%3 == 0 {
			fmt.Printf("s1 len:%d,cap%d,value:%v \n", len(s1), cap(s1), s1)
			println(s1)
		}
	}

	s2 := make([]int, 16)     // 声明一个长度和容量都为16的切片
	s3 := make([]int, 10, 32) // 10是长度,32是容量
	fmt.Println(len(s2), cap(s2))
	fmt.Println(s2)
	fmt.Println(len(s3), cap(s3))
	fmt.Println(s3)
	//s3[32] = 100 // 超出容量会报错,index out of range [33] with length 10
	s3 = append(s3, 100) // 超出长度,会自动扩容
	fmt.Println(s3)
}

func extendSlice() {
	arr := []int{1, 2, 3, 4, 5, 6, 8}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2, "为什么是[6 8]?")
	fmt.Println("s1的长度", len(s1), "容量", cap(s1))
	fmt.Println("s2的长度", len(s2), "容量", cap(s2))
	fmt.Println("从s1的容量可以看出,切片是引用类型,切片的容量是从切片的第一个元素到原切片的最后一个元素的长度,仅仅只是新建了一个指针,该指针指向了原数组的要切分的第一个元素")
	fmt.Println("所以s2其实是指向s1的第三个元素,也就是原切片的第五个元素,所以s2的长度是2,容量是2")
	println(arr[2:])
	println(s1)
	println(s1[3:])
	println(s2)
	fmt.Println("通过println方法打印Slice的ptr,可以验证以上说法,两个切片指向同一个地址且长度相同,区别在于容量")
}

func deleteElement() {
	s1 := []int{1, 2, 3, 4, 0, 0, 0, 0, 0}
	// 删除s1[3]
	s2 := append(s1[:3], s1[4:]...) // 因为append需要传入一个slice和添加的内容,所以无法直接传入slice内容,需要将后来的元素解压出来为int类型,使用...方法进行解压,类似于python的.value()
	fmt.Println(s2)
}
func deleteHeadAndTail() {
	s1 := []int{1, 2, 3, 4, 0, 0, 0, 0, 0}
	s2 := s1[1:]
	s3 := s2[:len(s2)-1]
	fmt.Println("s2", s2)
	fmt.Println("s3", s3)
	// 将s3复制到地址不同的变量中
	s4 := make([]int, len(s3))
	println(s3)
	copy(s4, s3)
	fmt.Println("s4", s4)
	println(s4)

}
