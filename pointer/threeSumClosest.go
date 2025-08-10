package pointer

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	diff := math.MaxInt
	res := 0

	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			third := len(nums) - 1
			for third > second {
				threeSum := nums[first] + nums[second] + nums[third]
				if threeSum == target {
					return target
				}
				if absInt(threeSum-target) < diff {
					diff = absInt(threeSum - target)
					res = threeSum
				}
				if threeSum <= target {
					break
				}
				third--
			}
		}
	}

	return res
}
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Show_threeSumClosest() {
	arr := []int{0, 0, 0}
	target := 1
	ans := threeSumClosest(arr, target)
	fmt.Printf("在数组%v中, 最接近:%d的三个数字的和为:%d\n", arr, target, ans)
}
