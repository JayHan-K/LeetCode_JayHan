package main

func main() {

	input := "pwwkew"
	answer := lengthOfLeongestSubstring(input)

	println(answer)
}

func stringInSlice(a string, list []string)bool{
	for _,b := range list{
		if b == a{
			return true
		}
	}
	return false
}


func lengthOfLeongestSubstring(s string) int {

	var answer = 0
	var longest = 0
	var lastItem = ""
	var nowLen = 0
	var nowItems []string
	for i:=0 ; i < len(s); i++{
		nowItem := s[i:i+1]
		if nowItem != lastItem{
			if(stringInSlice(nowItem, nowItems)){
				nowItems = nil
				nowItems = append(nowItems, nowItem)
				nowLen = 1

			}else{
				nowItems = append(nowItems, nowItem)
				nowLen++
			}
		}else{
			nowItems = nil
			nowItems = append(nowItems, nowItem)
			nowLen = 1
		}
		if longest < nowLen{
			longest = nowLen
		}
		lastItem = s[i:i+1]
	}

	answer = longest

	return answer
}