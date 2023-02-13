package main

import "fmt"

// IDatabaser 定义了数据库接口
type IDatabaser interface {
	Connect() error
	Disconnect() error
}

func HandleDB(dbs ...IDatabaser) {
	for _, db := range dbs {
		db.Connect()

	}
	for _, db := range dbs {
		db.Disconnect()
	}

}

func main() {
	mysql := Mysql{DBName: "localhost-Mysql"}
	redis := Redis{DBName: "localhost-Redis"}
	fmt.Println("正常不使用接口,开始连接数据库")
	mysql.Connect()
	redis.Connect()
	mysql.Disconnect()
	redis.Disconnect()

	// 当Mysql结构和Redis结构实现了Connect()和Disconnect()两个方法之后,两者的实例就是IDatabaser接口的一个实例,可以直接传递给HandleDB函数进行接口的操作
	// 面向对象编程,Go的多态
	fmt.Println("使用接口,开始连接数据库")
	HandleDB(&mysql, &redis) // 在这里可以直接将mysql(Mysql类型)和redis(Redis类型)赋值给接口变量dbs

}
