package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

func main() {

	l1 := &ListNode{Val: -1}
	inserVal(l1, 1)
	inserVal(l1, 8)
	//inserVal(l1, 3)

	l2 := &ListNode{Val: -1}
	inserVal(l2, 0)
	//inserVal(l2, 6)
	//inserVal(l2, 4)

	answer := addTwoNumbers(l1, l2)

	printListNode(answer)
}

func printListNode(list *ListNode){
	var temp = list
	for temp.Next != nil{
		fmt.Println(temp.Val)
		temp = temp.Next
	}
	fmt.Println(temp.Val)
}

func inserVal(list *ListNode, value int){
	var temp *ListNode
	temp = list

	if temp.Val == -1{
		list.Val = value
	}else{
		for temp.Next != nil{
			temp = temp.Next
		}
		new_node := &ListNode{Val:value}
		temp.Next = new_node
	}
}

func getValueList(list *ListNode)[] int{
	var values []int
	var temp = list
	for temp != nil{
		value := temp.Val
		values = append(values, value)
		temp = temp.Next
	}
	return values
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	answer := &ListNode{Val: -1}
	values1 := getValueList(l1)
	values2 := getValueList(l2)
	carry := 0
	maxLen := 0
	if len(values1) >= len(values2){
		maxLen = len(values1)
	}else{
		maxLen = len(values2)
	}


	for i:=0 ; i<maxLen ; i++{
		value1 := 0
		value2 := 0
		if len(values1) > i{
			value1 = values1[i]
		}
		if len(values2) > i {
			value2 = values2[i]
		}
		sum := value1 + value2
		if carry == 1{
			sum = sum + 1
			carry = 0
		}
		if sum >= 10{
			carry = 1
			sum = sum % 10
		}
		inserVal(answer, sum)
	}
	if(carry == 1){
		inserVal(answer, 1)
	}


	return answer
}
