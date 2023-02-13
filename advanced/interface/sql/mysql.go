package main

import "fmt"

type Mysql struct {
	DBName    string
	isConnect bool
}

func (mysql *Mysql) Connect() error {
	fmt.Print("正在连接数据库 => ", mysql.DBName)

	// 连接数据库
	mysql.isConnect = true

	if mysql.isConnect {
		fmt.Println("连接成功")
		return nil
	}
	return fmt.Errorf("连接失败")
}

func (mysql *Mysql) Disconnect() error {
	fmt.Println("正在断开连接数据库 => ", mysql.DBName)

	// 断开连接数据库
	mysql.isConnect = false

	return nil
}
