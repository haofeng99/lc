package tree

// https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
// 426-LCR 155. 将二叉搜索树转化为排序的双向链表
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{}
	var ans, pre *TreeNode
	for len(queue) > 0 && root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}

		root = queue[len(queue)-1]
		if ans == nil {
			ans = root
		}
		if pre != nil {
			pre.Right = root
			root.Left = pre
		}
		pre = root
		root = root.Right
		queue = queue[:len(queue)-1]
	}
	ans.Left = pre
	pre.Right = ans
	return ans
}
