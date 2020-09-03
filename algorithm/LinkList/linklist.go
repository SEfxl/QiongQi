package main

import (
	"fmt"
)
//定义节点
type LinkNode struct {
	data int
	next *LinkNode
}

//初始化头节点
func InitHeadNode() *LinkNode  {
	return &LinkNode{0,nil}
}

//头插法 head指针不断向前移动
func HeadInsert(arr [10]int) *LinkNode  {
	Length := len(arr)
	var head = InitHeadNode() //初始化头节点
	for i:=0; i<Length; i++ {
		node := &LinkNode{data:arr[i]}
		node.next = head
		head = node
	}
	return head
}

//尾插法 end指针不断后移,head保持不变
func TailInsert(arr [10]int) *LinkNode {
	Length := len(arr)
	var head = InitHeadNode() //初始化头节点
	end := head
	for i:=0; i<Length; i++ {
		node := &LinkNode{data:arr[i]}
		end.next = node
		end = node
	}
	return head
}

//遍历链表
func (Link *LinkNode) Show()  {
	p := Link
	for p != nil  {
		fmt.Printf("%d ",p.data)
		p = p.next
	}
	fmt.Println()
}

func main() {

	a := [...]int{1,2,3,4,5,6,7,8,9,10}
	p1 := HeadInsert(a)
	p1.Show()

	p2 := TailInsert(a)
	p2.Show()

}