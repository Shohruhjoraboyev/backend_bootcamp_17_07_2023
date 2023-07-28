package main

import "fmt"

func palindrome(arr []string) {
	polin:=[]string{}
	length:=len(arr)

	for i := 0; i < length; i++ {
		str:= arr[i]

			if str[0] == str[len(str)-1] {
				polin = append(polin, string(str))
			
			}
	
		
		
	}
	fmt.Println(polin)
}

func main() {
	arr:=[]string{"hello","ona","non","121","233","222","nun"}
	palindrome(arr)

}