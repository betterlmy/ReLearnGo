package main

type C1 interface {
	~int | ~int32 // 期望类型为int和int32
	M1()
}

type T1 struct {
	// struct T实现了接口C1所要求的方法
}

func (T1) M1() {

}

type T2 int // T2也满足C1所要求的方法

func (T2) M1() {

}

type T3 string

func foo[P C1](t P) {

}

func main() {
	//foo(*new(T1)) // T1类型未实现约束 'C1'，因为类型未包括在类型集('~int', '~int32')中
	// !!! new方法返回指针
	foo(*new(T2))
	//foo(*new(T3)) // T3类型未实现 'C1'，因为缺少某些方法: M1()
}
