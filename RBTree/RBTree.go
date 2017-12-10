package RBTree

import (
	"container/list"
	"fmt"
)

type RBTree struct {
	root *RBNode
}

//插入操作
func (rbTree *RBTree) Insert(entry Entryer) {
	if rbTree.root == nil {
		root := NewRBNode(entry)
		rbTree.insertCheck(root)
		return
	}
	rbTree.insertNode(rbTree.root, entry)
}

//查询节点
func (rbTree *RBTree) GetNode(entry Entryer) *RBNode {
	pNode := getNode(rbTree.root,entry)
	if pNode == nil {
		return pNode
	}
	result := new(RBNode)
	result.entry = pNode.entry
	return result
}

//删除操作
func (rbTree *RBTree) DeleteNode(entry Entryer) bool{
	query := getNode(rbTree.root,entry)
	if query == nil {
		return false
	}
	if query.left == nil || query.right == nil {
		rbTree.deleteOneNode(query)
	} else {
		//要删除的节点有两个子节点的时候 找到右子节点的最左子节点替换 删除该最左子节点变成删除一个节点的操作
		mostLeft := query.right
		for mostLeft.left != nil {
			mostLeft = mostLeft.left
		}
		query.entry = mostLeft.entry //替换
		mostLeft.color = RED
		rbTree.deleteOneNode(mostLeft)
	}
	return true
}

//中序遍历顺序输出
func (rbTree *RBTree) MidRec() {
	midRec(rbTree.root)
}

//层序遍历输出
func (rbTree *RBTree) LevelTraversal() {
	l := list.New()
	l.PushBack(rbTree.root)
	levelTraversal(l)
}

//层序遍历
func levelTraversal(l *list.List) {
	e := l.Front()
	l.Remove(e)
	for e != nil {
		v := e.Value
		pNode := v.(*RBNode)
		fmt.Print(pNode.entry.GetValue())
		fmt.Print(" ")
		fmt.Println(pNode.color)
		if pNode.left != nil {
			l.PushBack(pNode.left)
		}
		if pNode.right != nil {
			l.PushBack(pNode.right)
		}
		e = l.Front()
		if e != nil {
			l.Remove(e)
		}
	}
}

//中序遍历
func midRec(pNode *RBNode) {
	if pNode != nil {
		midRec(pNode.left)
		fmt.Print(pNode.entry.GetValue())
		fmt.Print(" ")
		fmt.Println(pNode.color)
		midRec(pNode.right)
	}
}

//查询节点
func getNode(pNode *RBNode, entry Entryer) *RBNode {
	if pNode == nil {
		return nil
	}
	res := pNode.entry.Compare(entry)
	if res == 0 {
		return pNode
	} else if res == -1 {
		return getNode(pNode.left, entry)
	} else {
		return getNode(pNode.right, entry)
	}
}

//插入节点
func (rbTree *RBTree) insertNode(pNode *RBNode, entry Entryer) {
	res := pNode.entry.Compare(entry)
	if res != 1 {
		if pNode.left != nil {
			rbTree.insertNode(pNode.left, entry)
		} else {
			temp := NewRBNode(entry)
			temp.parent = pNode
			pNode.left = temp
			rbTree.insertCheck(temp)
		}
	} else {
		if pNode.right != nil {
			rbTree.insertNode(pNode.right, entry)
		} else {
			temp := NewRBNode(entry)
			temp.parent = pNode
			pNode.right = temp
			rbTree.insertCheck(temp)
		}
	}
}

//检查插入
func (rbTree *RBTree) insertCheck(pNode *RBNode) {
	parent := pNode.parent
	if parent == nil {
		pNode.color = BLACK
		rbTree.root = pNode
		return
	}
	//父亲节点为红色则要处理
	if parent.color == RED {
		uncle := pNode.getUncle()
		if uncle != nil && uncle.color == RED {
			uncle.color = BLACK
			parent.color = BLACK
			parent.parent.color = RED
			rbTree.insertCheck(parent.parent)
		} else {
			grandParent := pNode.getGrandParent()
			if grandParent.left == parent {
				if parent.right == pNode {
					parent.leftRotate() //父节点先左旋
				}
				if root := grandParent.rightRotate(); root != nil {
					rbTree.root = root
				}
			} else {
				if parent.left == pNode {
					parent.rightRotate() //父节点先右旋
				}
				if root := grandParent.leftRotate(); root != nil {
					rbTree.root = root
				}
			}
			grandParent.color = RED
			parent.color = BLACK
		}
	}
}

//删除一个节点或没有节点的节点
func (rbTree *RBTree) deleteOneNode(rbNode *RBNode){
	var child *RBNode
	if rbNode.left == nil {
		child = rbNode.right
	} else {
		child = rbNode.left
	}
	parent := rbNode.parent
	if parent == nil {
		if child == nil {
			rbTree.root = nil
		} else {
			child.color = BLACK
			child.parent = nil
			rbTree.root = child
		}
		rbNode = nil
		return
	}
	if rbNode.color == RED {
		if parent.left == rbNode {
			parent.left = child
		} else {
			parent.right = child
		}
		if child != nil {
			child.parent = parent
		}
		rbNode = nil
		return
	}
	if child == nil{
		child = NewRBNode(Entryer{})
		child.parent = parent
		if parent.left == rbNode {
			parent.left = child
			//删除检查
			parent.left = nil
		} else {
			parent.right = child
			//删除检查
			parent.right = nil
		}
		child = nil
		rbNode = nil
		return
	}else {
		if parent.left == rbNode {
			parent.left = child
		} else {
			parent.right = child
		}
		child.parent = parent
		if child.color == RED {
			child.color = BLACK
			rbNode = nil
			return
		}
		//删除检查
		rbNode = nil
		return
	}

}

//删除检查
func (rbTree *RBTree) deleteCheck(rbNode *RBNode){
	parent := rbNode.parent
	if parent == nil{
		rbNode.color = BLACK
		rbTree.root = rbNode
		return
	}
	brother := rbNode.getSibling()
	if brother.color == RED {
		if parent.left == brother {
			parent.rightRotate()
		} else {
			parent.leftRotate()
		}
		parent.color = RED
		brother.color = BLACK
		brother = rbNode.getSibling()
		parent = rbNode.parent
	}
	s1Color := BLACK
	s2Color := BLACK
	if brother.left != nil {
		s1Color = brother.left.color
	}
	if brother.right != nil {
		s2Color = brother.right.color
	}
	if !s1Color && !s2Color {
		if parent.color == RED {
			parent.color = BLACK
			brother.color = RED
			return
		}
		brother.color = RED
		rbTree.deleteCheck(parent)
	}
	if parent.left == rbNode && s1Color && !s2Color {
		brother.color = RED
		brother.left.color = BLACK
		brother.rightRotate()
	} else if parent.right == rbNode && !s1Color && s2Color{
		brother.color = RED
		brother.right.color = BLACK
		brother.leftRotate()
	}
	brother.color = parent.color
	parent.color = BLACK
	if parent.left == rbNode {
		brother.right.color = BLACK
		parent.leftRotate()
	} else {
		brother.left.color = BLACK
		parent.rightRotate()
	}
}