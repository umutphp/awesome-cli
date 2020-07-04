package node_test

import (
	"github.com/umutphp/awesome-cli/internal/package/node"

	"testing"
)

func TestGetName(t *testing.T) {
	n := node.New("Test", "", "")
	if n.GetName() != "Test" {
		t.Errorf("Name is incorrect, got: %s, want: %s.", n.GetName(), "Test")
	}
}

func TestGetFancyText(t *testing.T) {
	n := node.New("Name", "", "")
	if n.GetFancyText() != "Name" {
		t.Errorf("Fancy text is incorrect, got: %s, want: %s.", n.GetFancyText(), "Name")
	}

	n = node.New("Name", "", "Description")
	if n.GetFancyText() != "Name Description" {
		t.Errorf("Fancy text is incorrect, got: %s, want: %s.", n.GetFancyText(), "Name Description")
	}
}

func TestGetPWD(t *testing.T) {
	nnn := node.New("NNN", "url", "Description NNN")
	nn := node.New("NN", "url", "Description NN")
	n := node.New("N", "url", "Description N")

	n.AddChild(nn)
	nn.AddChild(nnn)

	pwd := n.GetPWD()

	if pwd[0] != "N" {
		t.Errorf("Root node is incorrect, got: %s, want: %s ", pwd[0], "N")
	}

	pwd = n.GetChildren()[0].GetPWD()

	if pwd[0] != "NN" {
		t.Errorf("Root node is incorrect, got: %s, want: %s ", pwd[0], "NN")
	}

	if pwd[1] != "N" {
		t.Errorf("Second node is incorrect, got: %s, want: %s ", pwd[1], "N")
	}
}

func TestFindChildByName(t *testing.T) {
	parent := node.New("Parent", "url", "Description Parent")
	child1 := node.New("Child1", "url", "Description Child1")
	child2 := node.New("Child2", "url", "Description Child2")
	child3 := node.New("Child3", "url", "Description Child3")
	child4 := node.New("Child4", "url", "Description Child4")

	parent.AddChild(child1)
	parent.AddChild(child2)
	parent.AddChild(child3)
	parent.AddChild(child4)

	result := parent.FindChildByName("NoName")
	if result != nil {
		t.Errorf("Search result is incorrect, got: %s, want: %s ", result.GetName(), "Nil")
	}

	result = parent.FindChildByName("Child1")
	if result == nil {
		t.Errorf("Search result is incorrect, got: %s, want: %s ", result.GetName(), "Nil")
	}

	result = parent.FindChildByName("Child2")
	if result == nil {
		t.Errorf("Search result is incorrect, got: %s, want: %s ", result.GetName(), "Nil")
	}

	result = parent.FindChildByName("Child3")
	if result == nil {
		t.Errorf("Search result is incorrect, got: %s, want: %s ", result.GetName(), "Nil")
	}

	result = parent.FindChildByName("Child4")
	if result == nil {
		t.Errorf("Search result is incorrect, got: %s, want: %s ", result.GetName(), "Nil")
	}
}
