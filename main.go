package main

import (
	"fmt"

	tp "github.com/DickyMPriangga/LeetCodeSolution/string/findSubstring"
	//lList "github.com/DickyMPriangga/LeetCodeSolution/linkedList/utilities"
)

func main() {
	test1 := "wordgoodgoodgoodbestword"
	test2 := []string{"word", "good", "best", "good"}
	fmt.Println(tp.FindSubstring(test1, test2))
}
