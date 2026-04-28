package dynamicprogramming

import (
	"testing"
)

// ======================== 刚好交易k次 ========================

func TestMaxProfitVII_basic(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		k      int
		want   int
	}{
		{
			name:   "k=1 标准case",
			prices: []int{7, 1, 5, 3, 6, 4},
			k:      1,
			want:   5, // buy1 sell6
		},
		{
			name:   "k=2 两笔交易",
			prices: []int{3, 2, 6, 5, 0, 3},
			k:      2,
			want:   7, // buy2 sell6 + buy0 sell3
		},
		{
			name:   "k=2 单调递增-必须拆成两笔",
			prices: []int{1, 2, 3, 4, 5},
			k:      2,
			want:   3, // buy1 sell3 + buy4 sell5 = 2+1 (同日不可买卖)
		},
		{
			name:   "k=1 单调递增",
			prices: []int{1, 2, 3, 4, 5},
			k:      1,
			want:   4, // buy1 sell5
		},
		{
			name:   "k=1 单调递减",
			prices: []int{5, 4, 3, 2, 1},
			k:      1,
			want:   -1, // 强制1笔交易，最小亏损 buy2 sell1
		},
		{
			name:   "k=2 单调递减",
			prices: []int{5, 4, 3, 2, 1},
			k:      2,
			want:   -2, // 强制2笔交易，每笔最小亏损-1
		},
		{
			name:   "k=3 震荡行情",
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			k:      3,
			want:   8, // buy3 sell5(+2) + buy0 sell3(+3) + buy1 sell4(+3)
		},
		{
			name:   "k=0 零次交易",
			prices: []int{1, 2, 3, 4, 5},
			k:      0,
			want:   0,
		},
		{
			name:   "单天价格",
			prices: []int{5},
			k:      1,
			want:   0,
		},
		{
			name:   "k=2 先跌后涨再跌再涨",
			prices: []int{6, 1, 3, 2, 4, 7},
			k:      2,
			want:   7, // buy1 sell3 + buy2 sell7 = 2+5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_full", func(t *testing.T) {
			got := maxProfit_VII(tt.prices, tt.k)
			if got != tt.want {
				t.Errorf("maxProfit_VII(%v, %d) = %d, want %d", tt.prices, tt.k, got, tt.want)
			}
		})
		t.Run(tt.name+"_od", func(t *testing.T) {
			got := maxProfit_VII_(tt.prices, tt.k)
			if got != tt.want {
				t.Errorf("maxProfit_VII_(%v, %d) = %d, want %d", tt.prices, tt.k, got, tt.want)
			}
		})
	}
}

// 验证压缩版和完整版一致性
func TestMaxProfitVII_consistency(t *testing.T) {
	testCases := []struct {
		prices []int
		k      int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 2},
		{[]int{3, 2, 6, 5, 0, 3}, 3},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{5, 4, 3, 2, 1}, 2},
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 4},
		{[]int{6, 1, 3, 2, 4, 7}, 3},
		{[]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, 4},
	}

	for i, tc := range testCases {
		full := maxProfit_VII(tc.prices, tc.k)
		od := maxProfit_VII_(tc.prices, tc.k)
		if full != od {
			t.Errorf("case %d: full=%d, od=%d, prices=%v, k=%d", i, full, od, tc.prices, tc.k)
		}
	}
}

// ======================== 最少交易k次 ========================

func TestMaxProfitVIII_basic(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		k      int
		want   int
	}{
		{
			name:   "k=1 至少1次-等于不限次数场景",
			prices: []int{7, 1, 5, 3, 6, 4},
			k:      1,
			want:   7, // buy1 sell5 + buy3 sell6 = 4+3
		},
		{
			name:   "k=2 至少2次-最优需2次而非1次",
			prices: []int{1, 2, 3, 4, 5},
			k:      2,
			want:   3, // 至少2次: buy1 sell3(+2) + buy4 sell5(+1) = 3 (不限次数的最优4仅需1次)
		},
		{
			name:   "k=0 至少0次=不限次数",
			prices: []int{7, 1, 5, 3, 6, 4},
			k:      0,
			want:   7, // 同maxProfit_II
		},
		{
			name:   "k=1 单调递增-至少1次",
			prices: []int{1, 2, 3, 4, 5},
			k:      1,
			want:   4, // buy1 sell5 = 4 (1次 ≥ 1次)
		},
		{
			name:   "k=2 至少2次-天然有2次盈利机会",
			prices: []int{3, 2, 6, 5, 0, 3},
			k:      2,
			want:   7, // buy2 sell6 + buy0 sell3 = 4+3
		},
		{
			name:   "k=1 单调递减-至少1次",
			prices: []int{5, 4, 3, 2, 1},
			k:      1,
			want:   -1, // 强制至少1笔交易，最小亏损 buy2 sell1
		},
		{
			name:   "k=3 至少3次-长数组可完成3次",
			prices: []int{1, 2, 3, 4, 5, 6, 7},
			k:      3,
			want:   4, // buy1 sell2 + buy3 sell4 + buy5 sell6 + buy6 sell7...DP收敛到4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_full", func(t *testing.T) {
			got := maxProfit_VIII_(tt.prices, tt.k)
			if got != tt.want {
				t.Errorf("maxProfit_VIII_(%v, %d) = %d, want %d", tt.prices, tt.k, got, tt.want)
			}
		})
		t.Run(tt.name+"_od", func(t *testing.T) {
			got := maxProfit_VIII_od(tt.prices, tt.k)
			if got != tt.want {
				t.Errorf("maxProfit_VIII_od(%v, %d) = %d, want %d", tt.prices, tt.k, got, tt.want)
			}
		})
	}
}

// 验证最少k次: maxProfit_VIII（枚举法）vs maxProfit_VIII_（直接DP）一致性
func TestMaxProfitVIII_consistency(t *testing.T) {
	testCases := []struct {
		prices []int
		k      int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 1},
		{[]int{3, 2, 6, 5, 0, 3}, 2},
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{6, 1, 3, 2, 4, 7}, 1},
		{[]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, 2},
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 2},
	}

	for i, tc := range testCases {
		enum := maxProfit_VIII(tc.prices, tc.k)     // 枚举刚好k~k+5
		direct := maxProfit_VIII_(tc.prices, tc.k)  // 直接DP
		od := maxProfit_VIII_od(tc.prices, tc.k)    // 状态压缩

		if enum != direct {
			t.Errorf("case %d: enum=%d, direct=%d, prices=%v, k=%d", i, enum, direct, tc.prices, tc.k)
		}
		if direct != od {
			t.Errorf("case %d: direct=%d, od=%d, prices=%v, k=%d", i, direct, od, tc.prices, tc.k)
		}
	}
}
