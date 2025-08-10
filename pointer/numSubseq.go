package pointer

import (
	"fmt"
	"sort"
)

func numSubseq(nums []int, target int) int {
	Mod := 1000000000 + 7
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	// temp为一定包含左边届元素的子序列的个数
	temp := make([]int, n)
	temp[0] = 1
	for i := 1; i < n; i++ {
		temp[i] = (temp[i-1] << 1) % Mod
	}
	l, r := 0, n-1
	for l <= r {
		// 一直在枚举左边界,因此需要求和
		if nums[l]+nums[r] <= target {
			ans = (ans + temp[r-l]) % Mod
			l++
		} else {
			r--
		}
	}
	return ans
}

func Show_numSubseq() {
	nums := []int{3, 5, 6, 7}
	target := 9
	fmt.Println(numSubseq(nums, target))

}

// 子序列结论
// 1.长度为n的序列,不为空的子序列个数为 1 << n - 1, 及 2^n - 1,  每个元素一共有两种情况,出现在子序列中和不出现在子序中,因此一共2^n种可能,最后减去全都不存在的1种情况
// 1.1 为什么本题在枚举到合适的左右边界时,不直接使用2^n-1,因此此时左右边界内的子序列不一定满足题目条件,我们只能确定最小左边界,然后基于最小左边界组成的子序列求和
// 2.求一个序列中,一定包含某一个元素(左边元素)的子序列的个数 2^(n-1)
// 2.1 针对上述问题,首先可以通过快速幂求得结果;
// 2.2 还可以通过上述过程,记录过程中的结果,然后求和

// 快速幂
// https://blog.csdn.net/qq_19782019/article/details/85621386
// 1.结论1 求余公式
// (a + b) % p = (a % p + b % p) % p
// (a - b) % p = (a % p - b % p ) % p
// (a * b) % p = (a % p * b % p) % p

// 2.快速幂公式
func fastExp(base, power int) int {
	result := 1
	for power > 0 {
		if power%2 == 0 {
			power = power / 2
			base = base * base
		} else {
			power = power - 1
			result = result * base
			power = power / 2
			base = base * base
		}
	}
	return result
}

func fastExp2(base, power int) int {
	result := 1
	for power > 0 {
		if power%2 == 1 {
			result = result * base
		}
		power = power / 2
		base = base * base
	}
	return result
}

func fastExp3(base, power int) int {
	result := 1
	for power > 0 {
		if power&1 == 1 {
			result = result * base
		}
		power = power >> 1
		base = base * base
	}
	return result
}
