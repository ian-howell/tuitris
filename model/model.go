package model

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
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
)

const FPS = 1.0 / time.Second

// The Model contains the game's state.
//
// There should only ever be a single instance of a Model.
type Model struct {
	CurrentScreen screen.Screen

	MainViewport viewport.Model

	SplashModel   splash.Model
	MainMenuModel mainmenu.Model
	OptionsModel  options.Model
	ResetModel    reset.Model
	PlayModel     play.Model
	PauseModel    pause.Model
	WinModel      win.Model
	LoseModel     lose.Model
	ErrorModel    iamerror.Model
}

func (m Model) Init() tea.Cmd {
	return doTick()
}

// Update updates the game state. This happens in a goroutine alongside the View function.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TickMsg:
		return m, doTick()
	case tea.KeyMsg:
		key := msg.String()
		if key == "ctrl+c" || key == "q" {
			return m, tea.Quit
		}
	}

	switch m.CurrentScreen {
	case screen.Splash:
		m.SplashModel, m.CurrentScreen = m.SplashModel.Update(msg)
	case screen.MainMenu:
		m.MainMenuModel, m.CurrentScreen = m.MainMenuModel.Update(msg)
	case screen.Options:
		m.OptionsModel, m.CurrentScreen = m.OptionsModel.Update(msg)
	case screen.Reset:
		m.ResetModel, m.CurrentScreen = m.ResetModel.Update(msg)
	case screen.Play:
		m.PlayModel, m.CurrentScreen = m.PlayModel.Update(msg)
	case screen.Pause:
		m.PauseModel, m.CurrentScreen = m.PauseModel.Update(msg)
	case screen.Win:
		m.WinModel, m.CurrentScreen = m.WinModel.Update(msg)
	case screen.Lose:
		m.LoseModel, m.CurrentScreen = m.LoseModel.Update(msg)
	case screen.Error:
		m.ErrorModel, m.CurrentScreen = m.ErrorModel.Update(msg)

	}

	if m.CurrentScreen == screen.Exit {
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) View() string {
	var s string
	switch m.CurrentScreen {
	case screen.Splash:
		s = m.SplashModel.View()
	case screen.MainMenu:
		s = m.MainMenuModel.View()
	case screen.Options:
		s = m.OptionsModel.View()
	case screen.Reset:
		s = m.ResetModel.View()
	case screen.Play:
		s = m.PlayModel.View()
	case screen.Pause:
		s = m.PauseModel.View()
	case screen.Win:
		s = m.WinModel.View()
	case screen.Lose:
		s = m.LoseModel.View()
	case screen.Exit:
		// Nothing to print while exiting
	case screen.Error:
		s = m.ErrorModel.View()

	}

	m.MainViewport.SetContent(s)
	return m.MainViewport.View()
}
