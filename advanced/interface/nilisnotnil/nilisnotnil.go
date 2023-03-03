package main

// printNilInterface nil接口变量
func printNilInterface() {
	var i interface{} // 空接口
	var err error     //非空接口
	println(i)
	println(err)
	println("i == nil", i == nil)
	println("err == nil", err == nil)
	/*
		(0x0,0x0)
		(0x0,0x0)
		i == nil true
		err == nil true
		i==err true

		可以看出 无论空接口或者非空接口 一旦变量值为nil,它的内部表示均为(0x0,0x0) 类型信息,数据值信息
	*/
}

func printEmptyInterface() {
	var eif1, eif2 interface{}
	var n, m int = 17, 18

	eif1 = n
	eif2 = m
	println("eif1:", eif1)              //(0x10428c3e0,0x14000044760)
	println("eif2:", eif2)              // (0x10428c3e0,0x14000044758)
	println("eif1==eif2", eif1 == eif2) //false

	eif2 = 17
	println("eif1:", eif1)              //(0x10428c3e0,0x14000044760)
	println("eif2:", eif2)              //  (0x10428c3e0,0x104286a28)
	println("eif1==eif2", eif1 == eif2) //true

	eif2 = int64(17)
	println("eif1:", eif1) //(0x10428c3e0,0x14000044760)

	println("eif2:", eif2) //	println("eif2:", eif2)

	println("eif1 = eif2:", eif1 == eif2)
	/*output
	eif1: (0x10428c3e0,0x14000044760)
	eif2: (0x10428c3e0,0x14000044758)
	eif1==eif2 false
	eif1: (0x10428c3e0,0x14000044760)
	eif2: (0x10428c3e0,0x104286a28)
	eif1==eif2 true
	eif1: (0x10428c3e0,0x14000044760)
	eif2: (0x10428c4a0,0x104286a28)
	eif1 = eif2: false

	*/
}
