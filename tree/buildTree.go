package tree

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solutions/
// 进阶: 前序+中序 输出 后序遍历 TODO
func buildTreePI(preorder []int, inorder []int) *TreeNode {
	nums := len(preorder)

	iMap := make(map[int]int)
	for i, v := range inorder {
		iMap[v] = i
	}

	var dfs func(int, int, int, int) *TreeNode
	dfs = func(pL, pR, iL, iR int) *TreeNode {
		if pL > pR || iL > iR {
			return nil
		}
		iIdx := iMap[preorder[pL]]
		lCnt := iIdx - iL
		node := &TreeNode{Val: preorder[pL]}
		node.Left = dfs(pL+1, pL+lCnt, iL, iIdx-1)
		node.Right = dfs(pL+lCnt+1, pR, iIdx+1, iR)
		return node
	}
	return dfs(0, nums-1, 0, nums-1)
}

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// 中序 后序
func buildTreeIP(inorder []int, postorder []int) *TreeNode {
	nums := len(inorder)

	iMap := make(map[int]int)
	for i, v := range inorder {
		iMap[v] = i
	}

	var dfs func(int, int, int, int) *TreeNode
	dfs = func(il, ir, pl, pr int) *TreeNode {
		if il > ir || pl > pr {
			return nil
		}
		root := &TreeNode{Val: postorder[pr]}

		iIdx := iMap[postorder[pr]]

		root.Left = dfs(il, iIdx-1, pl, pl+(iIdx-il)-1)
		root.Right = dfs(iIdx+1, ir, pl+(iIdx-il), pr-1)

		return root
	}

	return dfs(0, nums-1, 0, nums-1)
}

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
// 前序, 后序
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	nums := len(preorder)

	postMap := make(map[int]int)
	for i, v := range postorder {
		postMap[v] = i
	}

	var dfs func(int, int, int, int) *TreeNode
	dfs = func(prel, prer, posl, posr int) *TreeNode {
		if prel > prer || posl > posr {
			return nil
		}
		root := &TreeNode{Val: preorder[prel]}
		if prel == prer {
			return root
		}
		postIdx := postMap[preorder[prel+1]] // 因为这里是加1操作, 所以需要前面判断一下

		leftTreeCnt := postIdx - posl + 1

		root.Left = dfs(prel+1, prel+leftTreeCnt, posl, posl+leftTreeCnt-1)
		root.Right = dfs(prel+leftTreeCnt+1, prer, postIdx+1, posr-1)
		return root
	}
	return dfs(0, nums-1, 0, nums-1)
}
