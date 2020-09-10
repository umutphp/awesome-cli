package main

import (
	"fmt"
	"os"

	"github.com/umutphp/awesome-cli/internal/package/favourite"
	"github.com/umutphp/awesome-cli/internal/package/manager"
	"github.com/umutphp/awesome-cli/internal/package/prompter"
)

const (
	// CACHE_KEY is the name of the cache folder
	CACHE_KEY = "awesome"
	// VERSION of the cli
	VERSION = "0.5.0"
)

func main() {
	args := os.Args[1:]
	manager := manager.New()

	manager.Initialize()

	fmt.Println("aweome-cli Version", VERSION)

	if len(args) > 0 {
		Argumented(args, manager)
		return
	}

	Walk(manager)
}

func DisplayHelp() {
	fmt.Println("")
	fmt.Println("Options of awesome-cli:")
	fmt.Printf("%-2v%-10v%-10v\n", "", "help", "To print this screen.")
	fmt.Printf("%-2v%-10v%-10v\n", "", "random", "To go to a random awesome content.")
	fmt.Printf("%-2v%-10v%-10v\n", "", "surprise", "To go to a surprise awesome content according to your previos choices.")
	fmt.Printf("%-2v%-10v%-10v\n", "", "profile", "To see your previos choices.")
	fmt.Printf("%-2v%-10v%-10v\n", "", "reset", "To clean your choices to start from the beginning.")
	fmt.Println("")
	fmt.Println("")
}

func RandomRepo(man manager.Manager) {
	rpwd, url := prompter.Random(&man)

	DisplayRepoWithPath(url, rpwd)
}

func SurpriseRepo(man manager.Manager) {
	favourites := favourite.NewFromCache(CACHE_KEY)

	if len(favourites.GetChildren()) == 0 {
		RandomRepo(man)
	}

	category := favourites.GetRandom()
	subcategory := category.GetRandom()
	rpwd, url := prompter.Surprise(&man, category.GetName(), subcategory.GetName())

	DisplayRepoWithPath(url, rpwd)
}

func Reset(man manager.Manager) {
	favourites := favourite.New(CACHE_KEY)
	favourites.SaveCache()
	fmt.Println("The choice list has been cleared.")
}

func Profile(man manager.Manager) {
	favourites := favourite.NewFromCache(CACHE_KEY)
	fmt.Println("")
	fmt.Println("Your choices:")

	for _, category := range favourites.GetChildren() {
		fmt.Println("", category.GetName())

		for _, subcategory := range category.GetChildren() {
			fmt.Println("  ", subcategory.GetName())
		}
	}

	fmt.Println("")
}

func DisplayRepoWithPath(url string, path []string) {
	for _, str := range path {
		fmt.Println(str)
	}

	fmt.Println(url)

	prompter.OpenInBrowser(url)
}

func Argumented(param []string, man manager.Manager) {
	if param[0] == "random" {
		RandomRepo(man)
		return
	}

	if param[0] == "surprise" {
		SurpriseRepo(man)
		return
	}

	if param[0] == "help" {
		DisplayHelp()
		return
	}

	if param[0] == "reset" {
		Reset(man)
		return
	}

	if param[0] == "profile" {
		Profile(man)
		return
	}
}

func Walk(man manager.Manager) {
	cursor := man.Root
	i := 0
	favourites := favourite.NewFromCache(CACHE_KEY)
	firstsel := ""

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

	fmt.Println(cursor.GetURL())

	favourites.SaveCache()

	prompter.OpenInBrowser(cursor.GetURL())

	walking := prompter.PromptToContinue()
	if walking == "y" {
		Walk(man)
	}
}
