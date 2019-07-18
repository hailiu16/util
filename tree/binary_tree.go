package tree

import (
	"errors"
	"fmt"
)

type Node struct {
	Value       interface{}
	Left, Right *Node
}

func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) Set(v interface{}) error {
	if node == nil {
		return errors.New("setting value to nil.node")
	}
	node.Value = v

	return nil
}

func (node *Node) POrder() {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	node.Left.POrder()
	node.Right.POrder()
}

func (node *Node) MOrder() {
	if node == nil {
		return
	}
	node.Left.POrder()
	fmt.Println(node.Value)
	node.Right.POrder()
}

func (node *Node) BOrder() {
	if node == nil {
		return
	}
	node.Left.POrder()
	node.Right.POrder()
	fmt.Println(node.Value)
}

func (node *Node) LOrder() {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	node.Left.POrder()
	node.Right.POrder()

}
