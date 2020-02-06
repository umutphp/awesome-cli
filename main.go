package main

import (
	"fmt"

	"awesome/internal/package/manager"
	"awesome/internal/package/prompter"
)

func main() {
	manager := manager.New()

	manager.Initialize()

	cursor := manager.Root
	i      := 0

	for {
	    prompt := prompter.Create(cursor.Name, cursor)

		_, selected, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", selected)

		prompter.ExecuteSelection(selected, &manager)
		cursor = manager.GetPWD()
		
		i++
		
		if i > 2 {
			break
		}
	}
}