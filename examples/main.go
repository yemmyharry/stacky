package main

import (
	"fmt"


	"github.com/yemmyharry/stacky"

)

func main() {
	myStack:= stacks.Stack{}
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)
	fmt.Println(myStack)
	fmt.Println("---------")
	fmt.Println(myStack.Peek())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Peek())
	fmt.Println(myStack)
	fmt.Println(myStack.IsEmpty())
	myQueue:= stacks.Queue{}
	myQueue.Enqueue(12)
	myQueue.Enqueue(20)
	myQueue.Enqueue(50)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}