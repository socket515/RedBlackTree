package RBTree

const (
	RED   bool = true
	BLACK bool = false
)

type RBNode struct {
	entry               Entryer
	color               bool
	parent, left, right *RBNode
}

func NewRBNode(entry Entryer) *RBNode {
	rbNoe := &RBNode{
		entry:  entry,
		color:  RED,
		parent: nil,
		left:   nil,
		right:  nil,
	}
	return rbNoe
}

// getGrandParent() 获取父级节点的父级节点
func (rbNode *RBNode) getGrandParent() *RBNode {
	parent := rbNode.parent
	if parent != nil {
		return parent.parent
	} else {
		return nil
	}
}

// getSibling() 获取兄弟节点
func (rbNode *RBNode) getSibling() *RBNode {
	parent := rbNode.parent
	if parent != nil {
		if rbNode == parent.left {
			return parent.right
		} else {
			return parent.left
		}
	} else {
		return nil
	}
}

// GetUncle() 父节点的兄弟节点
func (rbNode *RBNode) getUncle() *RBNode {
	parent := rbNode.parent
	if parent != nil {
		return parent.getSibling()
	} else {
		return nil
	}
}

//左旋参数为旋转轴的节点 若根节点变动返回根节点
func (rbNode *RBNode) leftRotate() *RBNode {
	var root *RBNode
	if rbNode == nil {
		return root
	}
	if rbNode.right == nil {
		return root
	}
	parent := rbNode.parent
	var isLeft bool
	if parent != nil {
		isLeft = parent.left == rbNode
	}
	grandson := rbNode.right.left
	if rbNode.right.left != nil {
		rbNode.right.left.parent = rbNode
	}
	rbNode.right.left = rbNode
	rbNode.parent = rbNode.right
	rbNode.right = grandson
	// 判断是否换了根节点
	if parent == nil {
		rbNode.parent.parent = nil
		root = rbNode.parent
	} else {
		if isLeft {
			parent.left = rbNode.parent
		} else {
			parent.right = rbNode.parent
		}
		rbNode.parent.parent = parent
	}
	return root
}

//右旋参数为旋转轴的节点 若根节点变动返回根节点
func (rbNode *RBNode) rightRotate() *RBNode {
	var root *RBNode
	if rbNode == nil {
		return root
	}
	if rbNode.left == nil {
		return root
	}
	parent := rbNode.parent
	var isLeft bool
	if parent != nil {
		isLeft = parent.left == rbNode
	}
	grandson := rbNode.left.right
	if grandson != nil {
		grandson.parent = rbNode
	}
	rbNode.left.right = rbNode
	rbNode.parent = rbNode.left
	rbNode.left = grandson
	// 判断是否换了根节点
	if parent == nil {
		rbNode.parent.parent = nil
		root = rbNode.parent
	} else {
		if isLeft {
			parent.left = rbNode.parent
		} else {
			parent.right = rbNode.parent
		}
		rbNode.parent.parent = parent
	}
	return root
}
