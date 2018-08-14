package main

func Sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func SumAll(noOfSlices ...[]int) []int {
	noOf := len(noOfSlices)
	var sumAll make([]int, noOf)
	for i, nums := range noOfSlices {
		sumAll[i] = Sum(nums)
	}
	return sumAll
}
