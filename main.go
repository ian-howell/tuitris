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
	checkError(err)

	initialModel := Model{
		CurrentScreen: MenuScreen,
		Menus: map[Screen]ring.Ring[Choice]{
			MenuScreen: menuMenu,
		},
	}

	p := tea.NewProgram(
		initialModel,
		tea.WithAltScreen(),
	)

	_, err = p.Run()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}

}
