
# 终极解法
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/solutions/740596/5xing-dai-ma-gao-ding-suo-you-gu-piao-ma-j6zo/


1.最多买卖1次  dp[i]表示第i天卖出的所得利润
2.买卖不限次数 dp[i][0]表示第i天卖出, dp[i][1]表示第i天卖出
3.最多买卖2笔  dp[i][status] 表示第i天, 在状态status下, 最大收益 status有5种状态
4.最多买卖k笔  sell[i]和buy[i] 分别表示买i次和卖i次的最大收益
5.最少买卖k笔
6.恰好买卖k笔
7.买卖不限次数_每笔带手续费  以不限次数2作为基础, 在买入时-fee, 或者卖出-fee, 如果以卖出时减去手续费, 初始化不用减手续费
8.买卖不限次数_带操作冷冻期 无限次数, 操作带有冷冻期  以最多买卖k笔为基础, 所有依赖下标都饿外减1


快速总结版: