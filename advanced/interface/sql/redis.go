package main

import "fmt"

type Redis struct {
	DBName string
}

func (redis Redis) Connect() error {
	fmt.Println("正在连接数据库 => ", redis.DBName)

	// TODO 连接数据库

	return nil
}

func (redis Redis) Disconnect() error {
	fmt.Println("正在断开连接数据库 => ", redis.DBName)

	// TODO 断开连接数据库

	return nil
}
