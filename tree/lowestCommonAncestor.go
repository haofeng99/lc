package tree

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// 235. 二叉搜索树的最近公共祖先
// 如果存在最近祖先,则一定满足下面两个条件
// 1. 该节点r为p或者q中的一个,且另外一个p或者q为节点r的子节点
// 2. p和q分布在节点r的左右子树中(一定不会分布在左子树或者右子树,一定是两边分布)

// 解法 自底向上递归
// 如果当前节点是 null 或者是目标节点之一（p 或 q），直接返回当前节点。
// 递归左右子树：
// 左子树返回值为 l，右子树返回值为 r。
// 根据左右子树的返回值判断：
// 如果左子树返回 null，说明 p 和 q 都在右子树中，返回右子树的结果。
// 如果右子树返回 null，说明 p 和 q 都在左子树中，返回左子树的结果。
// 如果左右子树都不为 null，说明当前节点就是最近公共祖先，返回当前节点。

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val == root.Val || q.Val == root.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
