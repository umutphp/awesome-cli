package main

import (
	"fmt"

	"awesome/internal/package/manager"
	"awesome/internal/package/prompter"

	"github.com/pkg/browser"
)

const VERSION = "0.0.2"

func main() {
	manager := manager.New()

	manager.Initialize()

	cursor := manager.Root
	i      := 0

	fmt.Println("aweome-cli Version", VERSION)

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

		//fmt.Printf("You choose %q\n", cursor.GetName())

		i++
		// Awesome tree has only three level depth
		if i > 3 {
			break
		}
	}

	browser.OpenURL(cursor.GetURL())
}