package main

import (
	"fmt"
	"os"

	"github.com/ian-howell/tuitris/play"
	"github.com/ian-howell/tuitris/ring"
	"github.com/ian-howell/tuitris/screen"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	mainWidth  = 44
	mainHeight = 44
)

func main() {

	splashMenu, err := ring.New(screen.MainMenu)
	checkError(err)

	menuMenu, err := ring.New(screen.Init, screen.Options, screen.Exit)
	checkError(err)

	optionsMenu, err := ring.New(screen.MainMenu)
	checkError(err)

	initScreen, err := ring.New(screen.Play)
	checkError(err)

	playScreen, err := ring.New(screen.Pause, screen.Win, screen.Lose)
	checkError(err)

	pauseScreen, err := ring.New(screen.Init, screen.MainMenu, screen.Exit)
	checkError(err)

	winScreen, err := ring.New(screen.Init, screen.MainMenu, screen.Exit)
	checkError(err)

	loseScreen, err := ring.New(screen.Init, screen.MainMenu, screen.Exit)
	checkError(err)

	mvp := viewport.New(mainWidth, mainHeight)
	mvp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1)

	playModel, err := play.New()
	checkError(err)

	initialModel := Model{
		CurrentScreen: screen.Splash,
		Menus: map[screen.Screen]ring.Ring[screen.Screen]{
			screen.Splash:   splashMenu,
			screen.MainMenu: menuMenu,
			screen.Options:  optionsMenu,
			screen.Init:     initScreen,
			screen.Play:     playScreen,
			screen.Pause:    pauseScreen,
			screen.Win:      winScreen,
			screen.Lose:     loseScreen,
		},
		MainViewport: mvp,
		PlayModel:    playModel,
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
