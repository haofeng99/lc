package pointer

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for third := second + 1; third < len(nums); third++ {
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}

				four := len(nums) - 1
				for four > third && nums[first]+nums[second]+nums[third]+nums[four] > target {
					four--
				}
				if third == four {
					// 如果当third和four编译到的数据都大于targer,那就没必要继续枚举后续的third+1和four
					// 因为我们的nums已经排序了,后续的数据只会越来越大
					break
				}
				if nums[first]+nums[second]+nums[third]+nums[four] < target {
					continue
				}
				if nums[first]+nums[second]+nums[third]+nums[four] == target {
					res = append(res, []int{nums[first], nums[second], nums[third], nums[four]})
				}
			}
		}
	}
	return res
}

func Show_fourSum() {
	arr := []int{2, 2, 2, 2, 2}
	targer := 8
	res := fourSum(arr, targer)
	fmt.Printf("返回结果:%v\n", res)
}
