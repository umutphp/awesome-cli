package main

import (
	"fmt"
	"os"

	"github.com/umutphp/awesome-cli/internal/package/manager"
	"github.com/umutphp/awesome-cli/internal/package/prompter"
	"github.com/umutphp/awesome-cli/internal/package/favourite"

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
	cursor     := man.Root
	i          := 0
	favourites := favourite.NewFromCache("awesome")
	firstsel   := ""

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

		// First selection
		if i == 0 {
			firstsel = selected
			favourites.Add(favourite.New(selected))
		}

		// Second selection
		if i == 1 {
			f := favourites.GetChild(firstsel)
			f.Add(favourite.New(selected))
		}

		i++
		// Awesome tree has only three level depth
		if i > 3 {
			break
		}
	}

	fmt.Println(favourites)
	fmt.Println(cursor.GetURL())
	
	favourites.SaveCache()

	browser.OpenURL(cursor.GetURL())
}