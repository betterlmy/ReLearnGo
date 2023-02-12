package main

import (
	"fmt"
	"testing"
)

// TestFibonacciInRecursion 测试FibonacciInRecursion
func TestFibonacciInRecursion(t *testing.T) {
	n := 45
	result := FibonacciInRecursion(n)
	expectedResult := 1134903170
	if result != expectedResult {
		t.Errorf("FibonacciInRecursion(%d) = %d, expectedResult = %d", n, result, expectedResult)
	}
}

func TestFibonacciInFor(t *testing.T) {
	n := 45
	result := FibonacciInFor(n)
	expectedResult := 1134903170
	if result != expectedResult {
		t.Errorf("FibonacciInFor(%d) = %d, expectedResult = %d", n, result, expectedResult)
	}
}

func TestFibonacciInSwitch(t *testing.T) {
	n := 45
	result := FibonacciInSwitch(n)
	expectedResult := 1134903170
	if result != expectedResult {
		t.Errorf("FibonacciInSwitch(%d) = %d, expectedResult = %d", n, result, expectedResult)
	}
}

func main() {
	//测试1-10的斐波那契数列

	//for i := 1; i < 10; i++ {
	//	fmt.Printf("斐波那契数列第%d项的值是:%d\n", i, FibonacciInFor(i))
	//	fmt.Printf("斐波那契数列第%d项的值是:%d\n", i, FibonacciInRecursion(i))
	//	fmt.Printf("斐波那契数列第%d项的值是:%d\n", i, FibonacciInSwitch(i))
	//}
	fmt.Println("The result of TestFibonacciInRecursion is:")
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, nil, nil, nil)
}
