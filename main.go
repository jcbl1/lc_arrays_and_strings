package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type IntervalsToBeSorted [][]int

func (its IntervalsToBeSorted) Len() int           { return len(its) }
func (its IntervalsToBeSorted) Less(i, j int) bool { return its[i][0] < its[j][0] }
func (its IntervalsToBeSorted) Swap(i, j int) {
	its[i], its[j] = its[j], its[i]
}

func main() {
	nums := []int{1, 1, 2}
	j := removeDuplicates(nums)
	fmt.Println(j)
	fmt.Println(nums[:j])
}

func pivotIndex(nums []int) int {
	getSum := func(ns []int) int {
		sum := 0
		for _, v := range ns {
			sum += v
		}
		return sum
	}
	total := getSum(nums)
	for i, v := range nums {
		left := getSum(nums[:i])
		if left == total-left-v {
			return i
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := right / 2
	for left <= right {
		if nums[mid] >= target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
			mid = (right-left)/2 + left
		} else {
			left = mid
			mid = (right-left)/2 + left
		}
		time.Sleep(time.Second)
		fmt.Println(left, right, mid)
	}
	return mid
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	if intervals[0][1] < intervals[1][0] {
		return append(intervals[:1], merge(intervals[1:])...)
	} else if intervals[0][1] < intervals[1][1] {
		return merge(append([][]int{{intervals[0][0], intervals[1][1]}}, intervals[2:]...))
	} else {
		return merge(append(intervals[:1], intervals[2:]...))
	}
}

func merge2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		lenMerged := len(merged)
		if intervals[i][0] > merged[lenMerged-1][1] {
			merged = append(merged, intervals[i])
		} else {
			if merged[lenMerged-1][1] < intervals[i][1] {
				merged[lenMerged-1][1] = intervals[i][1]
			}
		}
	}

	return merged
}

func rotate(matrix [][]int) {
	n := len(matrix)
	copied := make([][]int, n)
	for i := range copied {
		copied[i] = make([]int, n)
		for j := range copied[i] {
			copied[i][j] = matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = copied[n-j-1][i]
		}
	}
	fmt.Println(matrix)
}

func rotateImproved(matrix [][]int) {
	n := len(matrix)
	if n%2 == 0 {
		for i := 0; i < n/2; i++ {
			for j := 0; j < n/2; j++ {
				tmp := matrix[i][j]
				matrix[i][j] = matrix[n-j-1][i]
				matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
				matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
				matrix[j][n-i-1] = tmp
			}
		}
	} else {
		for i := 0; i < n/2; i++ {
			for j := 0; j < n/2+1; j++ {
				matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] = matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
			}
		}
	}
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	var row []int
	var col []int
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				exist := false
				for k := range row {
					if row[k] == i {
						exist = true
					}
				}
				if !exist {
					row = append(row, i)
				}
				exist = false
				for k := range col {
					if col[k] == j {
						exist = true
					}
				}
				if !exist {
					col = append(col, j)
				}
			}
		}
	}
	fmt.Println(row, col)

	for _, v := range row {
		for j := range matrix[v] {
			matrix[v][j] = 0
		}
	}
	for _, v := range col {
		for i := range matrix {
			matrix[i][v] = 0
		}
	}
	fmt.Println(matrix)
}

func setZeroes2(matrix [][]int) {
	n := len(matrix)
	copied := make([][]int, n)
	for i := range copied {
		copied[i] = make([]int, n)
		for j := range copied[i] {
			copied[i][j] = matrix[i][j]
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				for j1 := range copied[i] {
					copied[i][j1] = 0
				}
				for i1 := range copied {
					copied[i1][j] = 0
				}
			}
		}
	}
	copy(matrix, copied)
}

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])
	slice := []int{}
	dir := false
	for i, j := 0, 0; i < m && j < n; {
		slice = append(slice, mat[i][j])
		if i == 0 && (i+j)%2 == 0 && !dir && j != n-1 {
			j++
			dir = true
			continue
		} else if j == n-1 && (i+j)%2 == 0 && !dir {
			i++
			dir = true
			continue
		} else if j == 0 && (i+j)%2 == 1 && dir && i != m-1 {
			i++
			dir = false
			continue
		} else if i == m-1 && (i+j)%2 == 1 && dir {
			j++
			dir = false
			continue
		}
		switch dir {
		case false:
			i--
			j++
		case true:
			i++
			j--
		}
	}
	return slice
}

func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	n := len(strs[0])
	for i := 0; i < n; i++ {
		str := strs[0][:i+1]
		for j := 1; j < len(strs); j++ {
			if str != strs[j][:i+1] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func longestPalindrome(s string) string {
	pals := []string{}
	for i := range s {
		longest := string(s[i])
		for j := i + 1; j < len(s); j++ {
			diff := false
			for k := i; k <= j; k++ {
				if s[k] != s[j-k+i] {
					diff = true
					break
				}
			}
			if !diff {
				longest = s[i : j+1]
			}
		}
		pals = append(pals, longest)
	}
	sort.Slice(pals, func(i, j int) bool {
		return len(pals[i]) < len(pals[j])
	})
	return pals[len(pals)-1]
}

func longestPalindrome2(s string) string {
	var pal func(string) bool
	pal = func(s_ string) bool {
		if s_[0] != s_[len(s_)-1] {
			return false
		}
		if len(s_) <= 3 {
			return true
		}
		return pal(s_[1 : len(s_)-1])
	}

	pals := []string{}
	for i := range s {
		longest := string(s[i])
		for j := i + 1; j <= len(s); j++ {
			if pal(s[i:j]) {
				longest = s[i:j]
			}
		}
		pals = append(pals, longest)
	}

	sort.Slice(pals, func(i, j int) bool {
		return len(pals[i]) < len(pals[j])
	})

	return pals[len(pals)-1]
}

func reverseWords(s string) string {
	splited := strings.Split(s, " ")
	n := len(splited)
	res := ""
	for i := range splited {
		j := n - i - 1
		if splited[j] != "" {
			res += splited[j] + " "
		}
	}
	res = strings.TrimRight(res, " ")
	return res
}

func reverseString(s []byte) {
	n := len(s)
	j := n
	for i := 0; i < j; i++ {
		j = n - i - 1
		s[i], s[j] = s[j], s[i]
	}
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	i := 0
	j := n - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		switch {
		case sum == target:
			return []int{i + 1, j + 1}
		case sum > target:
			j--
		case sum < target:
			i++
		}
	}
	return nil
}

func removeElement(nums []int, val int) int {
	i2 := 0
	for i, v := range nums {
		if v == val {
			continue
		}
		nums[i2] = nums[i]
		i2++
	}
	return i2
}

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	counts := []int{}
	for _, v := range nums {
		if v == 1 {
			count++
			continue
		}
		if count != 0 {
			counts = append(counts, count)
			count = 0
		}
	}
	if count != 0 {
		counts = append(counts, count)
	}
	sort.Ints(counts)
	fmt.Println(counts)
	return counts[len(counts)-1]
}

func minSubArrayLen(target int, nums []int) int {
	sum := func(ns []int) int {
		res := 0
		for _, v := range ns {
			res += v
		}
		return res
	}
	min := len(nums)
	mins := []int{}
	for i := range nums {
		for j := len(nums) - 1; j >= i; j-- {
			if sum(nums[i:j+1]) < target {
				break
			}
			min = j + 1 - i
		}
		mins = append(mins, min)
	}
	sort.Ints(mins)
	return mins[0]
}

func fact(x int) int {
	if x <= 1 {
		return 1
	}
	return x * fact(x-1)
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = append(res[0], 1)
	for i := 1; i < numRows; i++ {
		res[i] = make([]int, i+1)
		for j := range res[i] {
			left, right := 0, 0
			if j > 0 {
				left = res[i-1][j-1]
			}
			if j < len(res[i-1]) {
				right = res[i-1][j]
			}
			res[i][j] = left + right
		}
	}
	return res
}

func getRow(rowIndex int) []int {
	// res:=make([]int,rowIndex+1)
	// for i:=range res{
	// 	res[i]=fact(rowIndex)/(fact(i)*fact(rowIndex-i))
	// }
	// return res

	triangle := generate(rowIndex)
	return triangle[rowIndex]
}

func reverseWord(s []byte) {
	n := len(s)
	// if n<2{
	// 	return
	// }
	// s[0],s[n-1]=s[n-1],s[0]
	// reverseWord(s[1:n-1])

	i := 0
	j := n
	for ; i < j; i++ {
		j = n - i - 1
		s[i], s[j] = s[j], s[i]
	}
}

func reverseWords3(s string) string {
	splited := strings.Split(s, " ")
	res := ""
	for i := range splited {
		wordSplited := []byte(splited[i])
		reverseWord(wordSplited)
		res += string(wordSplited) + " "
	}
	res = strings.TrimRight(res, " ")

	return res
}

func findMin(nums []int) int {
	n := len(nums)
	if nums[0] < nums[n-1] {
		return nums[0]
	}
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}
	return nums[0]
}

func removeDuplicates(nums []int) int {
	n := len(nums)

	j := 0
	for i := 1; i < n; i++ {
		if nums[i] == nums[j] {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	return j + 1
}

func moveZeroes(nums []int) {
	n := len(nums)
	res := make([]int, n)
	count := 0
	for _, v := range nums {
		if v != 0 {
			res[count] = v
			count++
		}
	}
	copy(nums, res)
}
