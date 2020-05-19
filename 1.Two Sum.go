package main

import "fmt"

func main() {
	nums :=[]int {11,7,2,15}
	var result = twoSum(nums, 9)
	fmt.Println("Result : " , result)

}

func twoSum(nums []int, target int)[]int {
	var answer []int

	for i:=0; i<len(nums)-1 ; i++{
		for j:=i+1 ; j<len(nums); j++{
			if nums[i] + nums[j] == target{
				answer = append(answer, i)
				answer = append(answer, j)
			}
		}
	}


	return answer
}
