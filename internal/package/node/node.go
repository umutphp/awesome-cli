package node

import (
	"fmt"
)

type Node struct {
    Name        string
    URL         string
    Description string
    Children    []Node
    Parent      *Node
}

func New(name, url, description string,) Node {
	return Node{
		Name:        name,
		URL:         url,
		Description: description,
		Children:    []Node{},
		Parent:      nil,
	}
}

func (n *Node) Display() {
	fmt.Println(n.Name, n.URL, n.Description)
} 

func (n *Node) GetName() string {
	return n.Name
}

func (n *Node) GetURL() string {
	return n.URL
}

func (n *Node) GetFancyText() string {
	fancy := n.GetName()

	if n.GetDescription() != "" {
		fancy = fancy + " " + n.GetDescription()
	}

	return fancy
}

func (n *Node) GetDescription() string {
	return n.Description
}

func (n *Node) GetParent() *Node {
	return n.Parent
}

func (n *Node) GetChildren() []Node {
	return n.Children
}

func (n *Node) SetChildren(arr []Node) {
	n.Children = arr
}

func (n *Node) GetPWD() []string {
	pwd   := []string{}
	point := n
	for {
		if (point.GetParent() == nil) {
			break
		}

		pwd = append(pwd, point.GetName())
		point = point.GetParent()
	}

	return pwd
}

func (n *Node) AddChild(child Node) {
	n.Children = append(n.Children, child)
}

func (n *Node) FindChildByName(name string) *Node {
	for _,child := range n.Children {
		if child.Name == name {
			return &child
		}
	}

	return nil
} 