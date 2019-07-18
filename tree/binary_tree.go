package tree

import (
	"errors"
	"fmt"
)

type Node struct {
	Value       interface{}
	left, right *Node
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

