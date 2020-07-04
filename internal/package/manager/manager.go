package manager

import (
	"fmt"

	"github.com/umutphp/awesome-cli/internal/package/fetcher"
	"github.com/umutphp/awesome-cli/internal/package/node"
	"github.com/umutphp/awesome-cli/internal/package/parser"
)

type Manager struct {
	Root          *node.Node
	PWD           *node.Node
	History       []Command
	ValidCommands []Command
}

type Command struct {
	Text string
}

func New() Manager {
	return Manager{
		Root:          nil,
		PWD:           nil,
		History:       []Command{},
		ValidCommands: []Command{},
	}
}

func (m *Manager) Execute(command Command) {
	command.Execute(m)
}

func (m *Manager) SetPWD(n *node.Node) {
	if len(n.GetChildren()) == 0 {
		fecthed, err := fetcher.FetchAwsomeRepo(n.GetURL())

		if err != nil {
			panic(err)
		}

		temp := parser.ParseIndex(fecthed)

		n.SetChildren(temp.GetChildren())
	}

	m.PWD = n
}

func (m *Manager) GetPWD() *node.Node {
	return m.PWD
}

func (m *Manager) GoBack() {
	if m.PWD.Parent != nil {
		m.SetPWD(m.PWD.Parent)
	}
}

func (m *Manager) Initialize() {
	fecthed, err := fetcher.FetchAwsomeRootRepo()

	if err != nil {
		panic(err)
	}

	root := parser.ParseIndex(fecthed)
	root.Name = "Awesome"
	m.Root = &root
	m.PWD = m.Root
}

func (c *Command) Execute(m *Manager) {
	switch c.Text {
	case "ls":
		fmt.Println(m.PWD.GetName())
	default:
		fmt.Println("Invalid command.")
	}
}
