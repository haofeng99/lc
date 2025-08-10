package tree

// https://leetcode.cn/problems/sum-root-to-leaf-numbers/
var ans int

func sumNumbers(root *TreeNode) int {
	ans = 0
	dfs_sumNumbers(root, 0)
	return ans
}

// 提前判断是否为子节点,否则会计算两遍
func dfs_sumNumbers(node *TreeNode, curVal int) {
	if node == nil {
		return
	}
	curVal = curVal*10 + node.Val
	if node.Left == nil && node.Right == nil {
		ans = ans + curVal
	}
	dfs_sumNumbers(node.Left, curVal)
	dfs_sumNumbers(node.Right, curVal)
}

// bfs思路:
// 1.原地修改:当遍历到某个节点时,如果不位叶子节点,则分别修改其子节点的值;如果为叶子节点,则直接求和;
// 2.不允许原地修改, 则额外开辟两个队列,分别存放节点以及对应的值, 需要注意的时,需要同时push和popl两个队列
func sumNumbers_bfs(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	nodeList := []*TreeNode{root}
	valList := []int{root.Val}

	for len(nodeList) != 0 {
		length := len(nodeList)
		for length > 0 {
			node := nodeList[0]
			if node.Left == nil && node.Right == nil {
				ans += valList[0]
			} else {
				if node.Left != nil {
					nodeList = append(nodeList, node.Left)
					valList = append(valList, valList[0]*10+node.Left.Val) // 新值的push
				}
				if node.Right != nil {
					nodeList = append(nodeList, node.Right)
					valList = append(valList, valList[0]*10+node.Right.Val)
				}
			}
			nodeList = nodeList[1:]
			valList = valList[1:]
			length--
		}
	}
	return ans
}
