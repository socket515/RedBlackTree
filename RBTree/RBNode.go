package RBTree



const (
	RED bool = true
	BLACK bool = false
)

type  RBNode struct {
	entry Entryer
	color bool
	parent,left,right *RBNode
}

func NewRBNode(entry Entryer) *RBNode {
	rbNoe := &RBNode{
		entry:entry,
		color:RED,
		parent:nil,
		left:nil,
		right:nil,
	}
	return rbNoe
}

// getGrandParent() 获取父级节点的父级节点
func(rbNode *RBNode) getGrandParent() *RBNode {
	parent := rbNode.parent
	if parent != nil {
		return parent.parent
	} else {
		return nil
	}
}
// getSibling() 获取兄弟节点
func(rbNode *RBNode) getSibling() *RBNode {
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
func(rbNode *RBNode) getUncle() *RBNode {
	parent := rbNode.parent
	if parent != nil {
		return parent.getSibling()
	} else {
		return nil
	}
}

//左旋参数为旋转轴的节点 若根节点变动返回根节点
func(rbNode *RBNode) leftRotate()*RBNode {
	right := rbNode.right
	if right == nil {
		return nil
	}
	grandSon := right.left
	right.left = rbNode
	rbNode.right = grandSon
	parent := rbNode.parent
	if parent != nil {
		if parent.left == rbNode {
			parent.left = right
		} else {
			parent.right = right
		}
		right.parent = parent
		return nil
	} else {
		right.parent = parent
		return right
	}
}

//右旋参数为旋转轴的节点 若根节点变动返回根节点
func(rbNode *RBNode) rightRotate()*RBNode {
	left := rbNode.left
	if left == nil {
		return nil
	}
	grandSon := left.right
	left.right = rbNode
	rbNode.left = grandSon
	parent := rbNode.parent
	if parent != nil {
		if parent.left == rbNode {
			parent.left = left
		} else {
			parent.right = left
		}
		left.parent = parent
		return nil
	} else {
		left.parent = parent
		return left
	}
}