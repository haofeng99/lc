package tree

import (
	"strconv"
	"strings"
)

/**
 * @param root:
 * @return: nothing
 */
func serialize(root *NodeNSub) string {
	if root == nil {
		return "[]"
	}

	var res []string
	var dfs func(*NodeNSub)
	dfs = func(node *NodeNSub) {
		if node == nil {
			// res = append(res, "nil")
			return
		}
		res = append(res, strconv.Itoa(node.val))
		res = append(res, strconv.Itoa(len(node.Children)))
		for _, sn := range node.Children {
			dfs(sn)
		}
	}
	dfs(root)
	return "[" + strings.Join(res, ",") + "]"
}

/**
 * @param data:
 * @return: nothing
 */
func deserialize(data string) *NodeNSub {
	if data == "[]" {
		return nil
	}
	data = data[1 : len(data)-1]
	dataArr := strings.Split(data, ",")

	idx := 0

	var dfs func() *NodeNSub
	dfs = func() *NodeNSub {
		val, _ := strconv.Atoi(dataArr[idx])
		idx++
		size, _ := strconv.Atoi(dataArr[idx])
		idx++

		node := &NodeNSub{val: val}
		queue := []*NodeNSub{}
		for i := 0; i < size; i++ {
			queue = append(queue, dfs())
		}
		node.Children = queue
		return node
	}
	return dfs()
}
