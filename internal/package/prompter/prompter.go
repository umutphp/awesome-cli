package prompter

import (
	"math/rand"
	"strings"
	"time"

	"github.com/toqueteos/webbrowser"
	"github.com/umutphp/awesome-cli/internal/package/fetcher"
	"github.com/umutphp/awesome-cli/internal/package/manager"
	"github.com/umutphp/awesome-cli/internal/package/node"

	"github.com/manifoldco/promptui"
)

func Create(title string, n *node.Node) promptui.Select {
	items := []string{}

	for _, child := range n.GetChildren() {
		items = append(items, child.GetName())
	}

	size := 5

	if len(items) > 10 {
		size = int(len(items) / 2)
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

func ToFavouriteString(child *node.Node) string {
	IconGood := promptui.Styler(promptui.FGGreen)("✔")
	return IconGood + " " + promptui.Styler(promptui.FGFaint)(strings.Trim(child.GetName(), " "))
}

func Surprise(m *manager.Manager, category string, subcategory string) ([]string, string) {
	cursor := m.Root
	list := []string{}

	child := cursor.FindChildByName(category)
	list = append(list, ToFavouriteString(child))

	child = child.FindChildByName(subcategory)
	list = append(list, ToFavouriteString(child))
	m.SetPWD(child)

	// Chose main category on sub awesome-list repository
	rand.Seed(time.Now().UnixNano())
	children := child.GetChildren()
	ind := rand.Intn(len(children))
	child = &children[ind]
	list = append(list, ToFavouriteString(child))

	m.SetPWD(child)

	// Select last child
	rand.Seed(time.Now().UnixNano())
	children = child.GetChildren()
	ind = rand.Intn(len(children))
	child = &children[ind]
	list = append(list, ToFavouriteString(child))

	if child.GetURL() == "" {
		return Surprise(m, category, subcategory)
	}

	if fetcher.IsUrl(child.GetURL()) == false {
		return Surprise(m, category, subcategory)
	}

	return list, child.GetURL()
}

func Random(m *manager.Manager) ([]string, string) {
	cursor := m.Root
	list := []string{}

	// Select main category
	rand.Seed(time.Now().UnixNano())
	children := cursor.GetChildren()

	ind := rand.Intn(len(children))
	child := &children[ind]
	list = append(list, ToFavouriteString(child))

	// Select sub awesome-list repository
	for {
		rand.Seed(time.Now().UnixNano())
		children = child.GetChildren()
		ind = rand.Intn(len(children))
		child = &children[ind]

		m.SetPWD(child)

		if len(child.GetChildren()) <= 1 {
			return Random(m)
		}

		list = append(list, ToFavouriteString(child))
		break
	}

	// Chose main category on sub awesome-list repository
	rand.Seed(time.Now().UnixNano())
	children = child.GetChildren()
	ind = rand.Intn(len(children))
	child = &children[ind]
	list = append(list, ToFavouriteString(child))

	m.SetPWD(child)

	// Select last child
	rand.Seed(time.Now().UnixNano())
	children = child.GetChildren()
	ind = rand.Intn(len(children))
	child = &children[ind]
	list = append(list, ToFavouriteString(child))

	if child.GetURL() == "" {
		return Random(m)
	}

	if fetcher.IsUrl(child.GetURL()) == false {
		return Random(m)
	}

	return list, child.GetURL()
}

func PromptToContinue() string {
	prompt := promptui.Prompt{
		Label:     "Continue walking in awesome lists",
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		return ""
	}

	return result
}

func OpenInBrowser(url string) {
	webbrowser.Open(url)
}
