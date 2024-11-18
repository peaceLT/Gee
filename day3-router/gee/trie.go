package gee

import "strings"

// 定义一个节点结构体，用于表示前缀树中的一个节点
type node struct {
	pattern  string  // 完整的路由路径，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点列表，例如 [doc, tutorial, intro]
	isWild   bool    // 是否为动态路由，part 含有 : 或 * 时为 true
}

// matchChild 返回第一个匹配成功的子节点，用于插入新节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		// 如果子节点的 part 与当前 part 相同，或者子节点是动态路由，则匹配成功
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 返回所有匹配成功的子节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		// 如果子节点的 part 与当前 part 相同，或者子节点是动态路由，则匹配成功
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 插入一个新的路由路径
func (n *node) insert(pattern string, parts []string, height int) {
	// 如果所有部分都已插入，则设置当前节点的 pattern
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// 如果没有匹配的子节点，则创建一个新的子节点
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 递归插入下一个部分
	child.insert(pattern, parts, height+1)
}

// search 查找与给定路径匹配的节点
func (n *node) search(parts []string, height int) *node {
	// 如果到达路径的末尾或遇到通配符节点，则返回当前节点
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		// 递归查找匹配的子节点
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
