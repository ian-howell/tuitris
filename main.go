package main

import (
	"fmt"
	"os"

	"github.com/ian-howell/tuitris/ring"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	menuMenu, err := ring.New(
		Choice{
			Name:       "Init",
			Cmd:        InitCmd,
			NextScreen: InitScreen,
		},
		Choice{
			Name:       "Exit",
			Cmd:        tea.Quit,
			NextScreen: ErrorScreen,
		},
	)
	if err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}

	initialModel := Model{
		CurrentScreen: MenuScreen,
		Menus: map[Screen]ring.Ring[Choice]{
			MenuScreen: menuMenu,
		},
	}
	if _, err := tea.NewProgram(initialModel, tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
