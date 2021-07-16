package main

import (
	"fmt"
	"sort"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(dp.longestPalindrome("babad"))
}

func maxArea(height []int) int {
	var result int = 0
	var waterLevel int = 1

	i := 0
	j := len(height) - 1

	for {
		for {
			if i == j || j < i {
				return result
			} else if height[i] < waterLevel {
				i++
			} else if height[j] < waterLevel {
				j--
			} else {
				break
			}
		}

		tempResult := (j - i) * waterLevel

		if tempResult > result {
			result = tempResult
		}

		waterLevel++
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	strInt := fmt.Sprint(x)

	var result string
	for _, v := range strInt {
		result = string(v) + result
	}

	return result == strInt
}

func convert(s string, numRows int) string {
	var result string

	if numRows == 1 {
		return s
	}

	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for j := i; j < len(s); j += (2 * (numRows - 1)) {
				result += string(s[j])
			}
		} else {
			inverse := true
			for j := i; j < len(s); {
				if inverse {
					result += string(s[j])
					j += 2 * (numRows - i - 1)
				} else {
					result += string(s[j])
					j += 2 * i
				}
				inverse = !inverse
			}
		}
	}

	return result
}

func mainFindMedianSortedArrays() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	fmt.Print(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums3 []int

	n3 := nums3[:]
	n3 = append(n3, nums1...)
	n3 = append(n3, nums2...)

	sort.Ints(n3)

	if len(n3)%2 == 0 {
		median1 := float64(n3[len(n3)/2])
		median2 := float64(n3[len(n3)/2-1])
		return (median1 + median2) / 2
	} else {
		return float64(n3[int(len(n3)/2)])
	}
}

func mainLengthOfLongestSubstring() {
	test := "abcabcbb"
	test2 := "bbbbb"
	test3 := "pwwkew"
	test4 := ""

	fmt.Print(lengthOfLongestSubstring(test))
	fmt.Print(lengthOfLongestSubstring(test2))
	fmt.Print(lengthOfLongestSubstring(test3))
	fmt.Print(lengthOfLongestSubstring(test4))
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	var resultString string
	var result = 0

	for i := 0; i < len(s); i++ {
		resultString = ""
		for j := i; j < len(s); j++ {
			if strings.Contains(resultString, string(s[j])) {
				break
			} else {
				resultString += string(s[j])
			}
		}

		if len(resultString) > result {
			result = len(resultString)
		}
	}

	return result
}

func mainAddTwoNumbers() {
	var l1 = ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	var l2 = ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

	addTwoNumbers(&l1, &l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result ListNode
	var p1, p2, pr *ListNode
	var remainder = 0

	p1 = l1
	p2 = l2
	pr = &result

	for {
		if p1 == nil {
			pr.Val = p2.Val + remainder
			p2 = p2.Next
		} else if p2 == nil {
			pr.Val = p1.Val + remainder
			p1 = p1.Next
		} else {
			pr.Val = p1.Val + p2.Val + remainder
			p1 = p1.Next
			p2 = p2.Next
		}

		if pr.Val >= 10 {
			remainder = 1
			pr.Val = pr.Val % 10
		} else {
			remainder = 0
		}

		if p1 == nil && p2 == nil {
			if remainder == 1 {
				pr.Next = &ListNode{Val: 1, Next: nil}
			}
			break
		} else {
			pr.Next = &ListNode{}
			pr = pr.Next
		}
	}

	return &result
}

func mainTwoSum() {
	var test = []int{2, 7, 11, 15}
	var testTarget = 9

	fmt.Print(twoSum(test, testTarget))
}

func twoSum(nums []int, target int) []int {
	var result = make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+nums[i] == target {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}

	return result
}
