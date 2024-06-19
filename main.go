package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	initialModel := Model{
		CurrentScreen: ScreenMenu,
		MenuMenu: Menu{
			Choices: []Choice{
				InitChoice{},
				ExitChoice{},
			},
		},
	}
	if _, err := tea.NewProgram(initialModel).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
