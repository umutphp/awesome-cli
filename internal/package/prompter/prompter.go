package prompter

import (
	"awesome/internal/package/node"
	"awesome/internal/package/manager"

	"github.com/manifoldco/promptui"
)

func Create(title string, n *node.Node) promptui.Select {
	items := []string{}

	for _,child := range n.GetChildren() {
		items = append(items, child.GetName())
	}

	size := 5

	if len(items) > 10 {
		size = int(len(items)/2)
	}

	return promptui.Select{
		Label: "Select from '" + title + "' list",
		Size:  size,
		Items: items,
	}
}

func ExecuteSelection(selected string, m *manager.Manager) {
	child := m.PWD.FindChildByName(selected)

	if child != nil {
		m.SetPWD(child)
	}
}