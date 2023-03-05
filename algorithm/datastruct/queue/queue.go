package queue

import (
	"errors"
	"fmt"
)

type intQueue struct {
	data []int64
	Len  int
	Head int // Head 指向队头的第一个数据
	Tail int // Tail 指向队尾最后一个空节点
}

// Dequeue 出队列
func (queue *intQueue) Dequeue() (int64, error) {
	if queue.Len == 0 {
		return -1, errors.New("队列长度为0,无法执行出列方法")
	}
	tmp := queue.data[queue.Head]
	queue.Len -= 1
	queue.Head += 1

	// 当有效元素过少时,重新分配
	if queue.Len <= len(queue.data)/2 {
		tmpQueueData := queue.data[queue.Head:queue.Tail]
		queue.data = tmpQueueData
		queue.Head = 0
		queue.Tail = queue.Len
	}

	fmt.Printf("当前队列剩余长度%d,容量%d\n", queue.Len, cap(queue.data))
	return tmp, nil
}

func (queue *intQueue) Enqueue(input int64) {
	queue.Len += 1
	queue.Tail += 1
	queue.data = append(queue.data, input)
	fmt.Printf("当前队列剩余长度%d\n", queue.Len)
}

func NewIntQueue(cap int) (*intQueue, error) {
	if cap <= 0 {
		return nil, errors.New("cap<=0")
	}
	return &intQueue{
		data: make([]int64, 0, cap),
		Len:  0,
		Head: 0,
		Tail: 0,
	}, nil
}

//func main() {
//	queue, _ := NewIntQueue(1)
//	for i := 0; i < 10; i++ {
//		queue.Enqueue(int64(i))
//	}
//	for i := 0; i < 10; i++ {
//		x, err := queue.Dequeue()
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(x)
//	}
//}
