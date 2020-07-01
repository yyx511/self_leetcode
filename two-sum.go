/*
两数之和（https://leetcode-cn.com/problems/two-sum/）

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func twoSum_1(nums []int, target int) []int {
   var i, j int
   len := len(nums)
   var res []int
   for i = 0; i < len; i++ {
	   for j = i + 1; j < len; j++ {
		   if nums[j] + nums[i] == target {
			   return append(res, i, j)
		   }
	   }
   }
   return res
}

func twoSum_2(nums []int, target int) []int {
	var i int
	var res []int
	var m = map[int]int{}
	len := len(nums)
	for i = 0; i < len; i++ {
		l := target - nums[i]
		if _, ok := m[l]; ok {
			return append(res, i, m[l])
		}
		m[nums[i]] = i
	}
	return res
}

func main() {
	var nums = []int{2, 7, 11, 15}
	res := twoSum_1(nums, 9)
	fmt.Printf("%v", res)

	res = twoSum_2(nums, 9)
	fmt.Printf("%v", res)
}

