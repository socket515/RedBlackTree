package RBTree


type Entryer interface {
	GetValue()interface{}
	Compare(Entryer)int
}
