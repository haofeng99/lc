package tree

// https://www.nowcoder.com/questionTerminal/9023a0c988684a53960365b889ceaf5e?
// 给定一个二叉树其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。注意，树中的结点不仅包含左右子结点，同时包含指向父结点的next指针
// 要求: 空间复杂度O(1), 时间复杂度O(n)
type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

// 1.存在右子树, 右子树最左边界即位下一个节点
// 2.不存在右子树, 找到其父节点, 如果父节点的左子树是当前节点, 则父节点就是下一个节点
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode == nil {
		return nil
	}
	if pNode.Right != nil {
		pNode = pNode.Right
		for pNode.Left != nil {
			pNode = pNode.Left
		}
		return pNode
	} else {
		pNodeParnet := pNode.Next
		for pNodeParnet != nil {
			if pNodeParnet.Left == pNode {
				return pNodeParnet
			}
			pNode = pNodeParnet
			pNodeParnet = pNodeParnet.Next
		}
	}
	return nil
}
