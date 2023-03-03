package main

import "fmt"

// QuackableAnimal 定义一个接口,只要实现了Quack方法的动物都可以当做这个接口的实例
type QuackableAnimal interface {
	Quack()
}

type Dog struct {
}

func (dog1 Dog) Quack() {
	fmt.Printf("有狗在叫! 汪!\n")
}

type Cat struct {
}

func (cat1 Cat) Quack() {
	fmt.Printf("有猫在叫! 喵!\n")
}

func (duck1 Duck) Quack() {
	fmt.Printf("有鸭在叫! 嘎!\n")
}

func (duck1 Duck) Run() {
	fmt.Printf("有鸭在跑! \n")
}

type Duck struct {
}

func AnAnimalQuackInAForest(animal QuackableAnimal) {
	animal.Quack()
}

func main() {
	cat1 := Cat{}
	duck1 := Duck{}
	dog1 := Dog{}

	AnAnimalQuackInAForest(cat1)
	AnAnimalQuackInAForest(duck1)
	AnAnimalQuackInAForest(dog1)

	/*
		这就体现了Go语言中接口的特性,同一个函数,只要传入一个接口即可,而不用去操心其传入的类别(不需要针对这个类别进行单独的操作),只要这个类能够实现Quack这个方法即可.
	*/

	// 也可以写成如下的形式
	animals := []QuackableAnimal{new(Duck), new(Dog), new(Cat)}
	for _, animal := range animals {
		AnAnimalQuackInAForest(animal)
	}
}
