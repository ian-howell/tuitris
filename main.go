package main

import (
	"fmt"
	"os"

	"github.com/ian-howell/tuitris/model"
	"github.com/ian-howell/tuitris/model/iamerror"
	"github.com/ian-howell/tuitris/model/lose"
	"github.com/ian-howell/tuitris/model/mainmenu"
	"github.com/ian-howell/tuitris/model/options"
	"github.com/ian-howell/tuitris/model/pause"
	"github.com/ian-howell/tuitris/model/play"
	"github.com/ian-howell/tuitris/model/reset"
	"github.com/ian-howell/tuitris/model/splash"
	"github.com/ian-howell/tuitris/model/win"
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

	initialModel := model.Model{
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
