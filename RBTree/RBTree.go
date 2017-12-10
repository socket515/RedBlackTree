package RBTree

import (
	"fmt"
	"container/list"
)

type RBTree struct {
	root *RBNode
}

//插入操作
func(rbTree *RBTree) Insert(entry Entryer) {
	if rbTree.root == nil {
		root := NewRBNode(entry)
		rbTree.insertCheck(root)
		return
	}
	rbTree.insertNode(rbTree.root,entry)
}

//查询节点
func(rbTree *RBTree) GetNode(pNode *RBNode,entry Entryer) *Entryer {
	return getNode(rbTree.root,entry)
}

//中序遍历顺序输出
func(rbTree *RBTree) MidRec(){
	midRec(rbTree.root)
}

//层序遍历输出
func(rbTree *RBTree) LevelTraversal(){
	l := list.New()
	l.PushBack(rbTree.root)
	levelTraversal(l)
}

//层序遍历
func  levelTraversal(l *list.List){
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
		if e!= nil {
			l.Remove(e)
		}
	}
}

//中序遍历
func midRec(pNode *RBNode){
	if pNode != nil {
		midRec(pNode.left)
		fmt.Print(pNode.entry.GetValue())
		fmt.Print(" ")
		fmt.Println(pNode.color)
		midRec(pNode.right)
	}
}

//查询节点
func getNode(pNode *RBNode,entry Entryer) *Entryer {
	if pNode == nil {
		return nil
	}
	res := pNode.entry.Compare(entry)
	if res == 0 {
		return &(pNode.entry)
	} else if res == -1 {
		return getNode(pNode.left,entry)
	} else {
		return getNode(pNode.right,entry)
	}
}

//插入节点
func(rbTree *RBTree) insertNode(pNode *RBNode,entry Entryer){
	res := pNode.entry.Compare(entry)
	if res != 1 {
		if pNode.left != nil {
			rbTree.insertNode(pNode.left,entry)
		} else {
			temp := NewRBNode(entry)
			temp.parent = pNode
			pNode.left = temp
			rbTree.insertCheck(temp)
		}
	} else {
		if pNode.right != nil {
			rbTree.insertNode(pNode.right,entry)
		} else {
			temp := NewRBNode(entry)
			temp.parent = pNode
			pNode.right = temp
			rbTree.insertCheck(temp)
		}
	}
}

//检查插入
func(rbTree *RBTree) insertCheck(pNode *RBNode){
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
					parent.leftRotate()//父节点先左旋
				}
				if root := grandParent.rightRotate(); root != nil {
					rbTree.root = root
				}
			} else {
				if parent.left == pNode {
					parent.rightRotate()//父节点先右旋
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