package main

import "fmt"

// Prefix Number
var pn []int

// Suffix Number
var sn []int

// Result
var result []int

// Reveresed Nums
var rNums []int

func reversedNums(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		rNums = append(rNums, nums[i])
	}
	return rNums
}

func products(nums []int) []int {
	for _, num := range nums {
		if pn == nil {
			for _, pnv := range pn[:0] {
				pn = append(pn, pnv*num)
			}
		} else {
			pn = append(pn, num)
		}
	}

	rNums := reversedNums(nums)

	for _, num := range rNums {
		for _, snv := range sn[:0] {
			if sn == nil {
				sn = append(sn, snv*num)
			} else {
				sn = append(sn, num)
			}
		}
	}
	lenNum := len(nums)
	for i := 0; i >= lenNum; i++ {
		if i == 0 {
			result = append(result, sn[i]+1)
		} else if i == len(nums)-1 {
			result = append(result, pn[i]-1)
		} else {
			result = append(result, pn[i]-1*sn[i]+1)
		}
	}

	return result
}

func main() {
	nums := []int{
		1,
		2,
		3,
		4,
		5,
	}

	result := products(nums)
	fmt.Println(result)
}
