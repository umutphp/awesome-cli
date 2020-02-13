package prompter

import (
	"math/rand"
    "time"
    "strings"
	
	"github.com/umutphp/awesome-cli/internal/package/node"
	"github.com/umutphp/awesome-cli/internal/package/manager"

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

	if size > 10 {
		size = 10
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

func Random(m *manager.Manager) ([]string, string) {
	IconGood := promptui.Styler(promptui.FGGreen)("✔")
	cursor   := m.Root
	list     := []string{}

	// Select main category
	rand.Seed(time.Now().UnixNano())
	children := cursor.GetChildren()

    ind      := rand.Intn(len(children))
    child    := &children[ind]
    list      = append(list, IconGood + " " + strings.Trim(child.GetName(), " "))

    // Select sub awesome-list repository
    for {
	    rand.Seed(time.Now().UnixNano())
	    children  = child.GetChildren()
	    ind       = rand.Intn(len(children))
	    child     = &children[ind]
	    list      = append(list, IconGood + " " + child.GetName())
		m.SetPWD(child)

		if len(child.GetChildren()) > 1 {
			break;
		}
	}

	// Chose main category on sub awesome-list repository
	rand.Seed(time.Now().UnixNano())
    children  = child.GetChildren()
    ind       = rand.Intn(len(children))
    child     = &children[ind]
    list      = append(list, IconGood + " " + child.GetName())

	m.SetPWD(child)

	// Select last child
	rand.Seed(time.Now().UnixNano())
    children  = child.GetChildren()
    ind       = rand.Intn(len(children))
    child     = &children[ind]
    list      = append(list, IconGood + " " + child.GetName())

    if child.GetURL() == "" {
    	return Random(m)
    }

	return list,child.GetURL()
}