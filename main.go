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

	fmt.Println("aweome-cli Version", VERSION)

	if len(args) > 0 && (args[0] == "random" || args[0] == "surprise") {
		RandomRepo(manager)
		return
	}

	if len(args) > 0 && args[0] == "help" {
		DisplayHelp()
		return
	}
		
	Walk(manager)
}

func DisplayHelp() {
	fmt.Println("")
	fmt.Println("Options of awesome-cli:")
	fmt.Printf("%-2v%-10v%-10v\n", "", "help", "To print this screen.")
	fmt.Printf("%-2v%-10v%-10v\n", "", "random", "To go to a random awesome content.")
	fmt.Printf("%-2v%-10v%-10v\n", "", "surprise", "Same with random option.")
}

func RandomRepo(man manager.Manager) {
	rpwd,url := prompter.Random(&man)

	for _, str := range rpwd {
		fmt.Println(str)
	}

	fmt.Println(url)
	
	browser.OpenURL(url)
}

func Walk(man manager.Manager) {
	cursor := man.Root
	i      := 0

	for {
	    prompt := prompter.Create(cursor.Name, cursor)

		_, selected, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		prompter.ExecuteSelection(selected, &man)

		// Where we are in the three
		cursor = man.GetPWD()

		i++
		// Awesome tree has only three level depth
		if i > 3 {
			break
		}
	}

	fmt.Println(cursor.GetURL())

	browser.OpenURL(cursor.GetURL())
}