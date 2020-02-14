package main

import (
	"fmt"
	"os"

	"github.com/umutphp/awesome-cli/internal/package/manager"
	"github.com/umutphp/awesome-cli/internal/package/prompter"

	"github.com/pkg/browser"
)

const VERSION = "0.2.0"

func main() {
    args    := os.Args[1:]
	manager := manager.New()

	manager.Initialize()

	cursor := manager.Root
	i      := 0

	fmt.Println("aweome-cli Version", VERSION)

	if len(args) > 0 && (args[0] == "random" || args[0] == "surprise") {
		RandomRepo(manager)
		return
	}
		
	for {
	    prompt := prompter.Create(cursor.Name, cursor)

		_, selected, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		prompter.ExecuteSelection(selected, &manager)

		// Where we are in the three
		cursor = manager.GetPWD()

		i++
		// Awesome tree has only three level depth
		if i > 3 {
			break
		}
	}

	browser.OpenURL(cursor.GetURL())
}

func RandomRepo(man manager.Manager) {
	rpwd,url := prompter.Random(&man)

	for _, str := range rpwd {
		fmt.Println(str)
	}

	browser.OpenURL(url)
}