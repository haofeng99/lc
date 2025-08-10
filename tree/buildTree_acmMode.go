package tree

// acm模式下 从array中创建一棵树

func buildTreeFromArr(arr []any) *TreeNode {
	return BuildTreeBFS(arr)
}

func BuildTreeBFS(data []any) *TreeNode {
	if len(data) == 0 || data[0] == nil {
		return nil
	}

	root := &TreeNode{Val: data[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(data) && len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// 左孩子
		if i < len(data) {
			if data[i] != nil {
				left := &TreeNode{Val: data[i].(int)}
				curr.Left = left
				queue = append(queue, left)
			}
			i++
		}

		// 右孩子
		if i < len(data) {
			if data[i] != nil {
				right := &TreeNode{Val: data[i].(int)}
				curr.Right = right
				queue = append(queue, right)
			}
			i++
		}
	}

	return root
}
