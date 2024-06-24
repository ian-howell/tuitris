package main

import (
	"fmt"
	"os"

	"github.com/ian-howell/tuitris/iamerror"
	"github.com/ian-howell/tuitris/lose"
	"github.com/ian-howell/tuitris/mainmenu"
	"github.com/ian-howell/tuitris/options"
	"github.com/ian-howell/tuitris/pause"
	"github.com/ian-howell/tuitris/play"
	"github.com/ian-howell/tuitris/reset"
	"github.com/ian-howell/tuitris/screen"
	"github.com/ian-howell/tuitris/splash"
	"github.com/ian-howell/tuitris/win"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	mainWidth  = 44
	mainHeight = 44
)

func main() {

	mvp := viewport.New(mainWidth, mainHeight)
	mvp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1)

	splashModel, err := splash.New()
	checkError(err)

	mainMenuModel, err := mainmenu.New()
	checkError(err)

	optionsModel, err := options.New()
	checkError(err)

	resetModel, err := reset.New()
	checkError(err)

	playModel, err := play.New()
	checkError(err)

	pauseModel, err := pause.New()
	checkError(err)

	winModel, err := win.New()
	checkError(err)

	loseModel, err := lose.New()
	checkError(err)

	errorModel, err := iamerror.New()
	checkError(err)

	initialModel := Model{
		CurrentScreen: screen.Splash,

		MainViewport: mvp,

		SplashModel:   splashModel,
		MainMenuModel: mainMenuModel,
		OptionsModel:  optionsModel,
		ResetModel:    resetModel,
		PlayModel:     playModel,
		PauseModel:    pauseModel,
		WinModel:      winModel,
		LoseModel:     loseModel,
		ErrorModel:    errorModel,
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
