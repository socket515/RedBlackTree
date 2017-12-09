package RBTree


type Entryer interface {
	SetValue(interface{})
	GetValue()interface{}
	Compare(Entryer)int
}
