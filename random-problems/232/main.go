package main

import "fmt"

type MyQueue struct {
    S1 []int
    S2 []int
}

func Constructor() MyQueue {
    return MyQueue{}
}

func (mq *MyQueue) Push(x int)  {
    mq.S1 = append(mq.S1 , x)
}

func (mq *MyQueue) Pop() int {
    if len(mq.S2) == 0 {
        for len(mq.S1) != 0 {
            index := len(mq.S1) - 1
            element := mq.S1[index]
            mq.S1 = mq.S1[:index]
            mq.S2 = append(mq.S2, element)
        }
    }
    
    index := len(mq.S2) - 1
    element := mq.S2[index]
    mq.S2 = mq.S2[:index]
    
    return element
}


func (mq *MyQueue) Peek() int {
    if len(mq.S2) == 0 {
        for len(mq.S1) != 0 {
            index := len(mq.S1) - 1
            element := mq.S1[index]
            mq.S1 = mq.S1[:index]
            mq.S2 = append(mq.S2, element)
        }
    }
    
    index := len(mq.S2) - 1
    element := mq.S2[index]
    
    return element
}


func (mq *MyQueue) Empty() bool {
    if len(mq.S1) == 0 && len(mq.S2) == 0 {
        return true
    }
    return false
}

func main(){
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Peek()
	param_3 := obj.Pop()
	param_4 := obj.Empty()
	param_5 := obj.Pop()
	param_6 := obj.Empty()
	fmt.Println(param_2, param_3, param_4, param_5, param_6)
}