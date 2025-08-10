package tree

import (
	"strconv"
	"strings"
)

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

type Codec2 struct {
}

func Constructor2() Codec2 {
	return Codec2{}
}

// Serializes a tree to a single string.
// 层序遍历结果
func (this *Codec2) serialize2(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	ans := "["
	nodeList := []*TreeNode{root}
	for len(nodeList) > 0 {
		node := nodeList[0]
		if node != nil {
			nodeList = append(nodeList, node.Left)
			nodeList = append(nodeList, node.Right)
			ans = ans + strconv.Itoa(node.Val) + ","
		} else {
			ans = ans + "nil" + ","
		}

		nodeList = nodeList[1:]
	}
	ans = ans[:len(ans)-1]
	return ans + "]"
}

// Deserializes your encoded data to tree.
// 层序遍历结果
func (this *Codec2) deserialize2(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	data = data[1 : len(data)-1]
	vals := strings.Split(data, ",")

	val, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: val}
	nodeList := []*TreeNode{root}

	cursor := 1

	for len(nodeList) > 0 {
		node := nodeList[0]
		valStrL := vals[cursor]
		if "nil" != valStrL {
			val, _ := strconv.Atoi(valStrL)
			node.Left = &TreeNode{Val: val}
			nodeList = append(nodeList, node.Left)
		}
		cursor++

		valStrR := vals[cursor]
		if "nil" != valStrR {
			val, _ := strconv.Atoi(valStrR)
			node.Right = &TreeNode{Val: val}
			nodeList = append(nodeList, node.Right)
		}
		cursor++

		nodeList = nodeList[1:]
	}
	return root
}
