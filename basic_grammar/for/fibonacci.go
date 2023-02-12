package main

// FibonacciInFor 使用For循环实现斐波那契数列
func FibonacciInFor(n int) int {
	// 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610
	if n == 1 || n == 2 {
		return 1
	} else {
		// 不断更新a和b,最后返回 a+b
		a, b := 1, 1
		for i := 3; i < n; i++ {
			a, b = b, a+b
		}
		return a + b
	}

}

// FibonacciInRecursion 使用递归实现斐波那契数列
func FibonacciInRecursion(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else if n <= 0 {
		return 0
	} else {
		return FibonacciInRecursion(n-2) + FibonacciInRecursion(n-1)
	}

}

// FibonacciInSwitch 使用递归实现斐波那契数列
func FibonacciInSwitch(n int) int {
	switch n {
	case 1, 2:
		return 1
	case 0:
		return 0
	default:
		return FibonacciInRecursion(n-2) + FibonacciInRecursion(n-1)
	}

}
