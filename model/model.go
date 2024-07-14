package model

import (
	"errors"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ian-howell/tuitris/model/iamerror"
	"github.com/ian-howell/tuitris/model/lose"
	"github.com/ian-howell/tuitris/model/mainmenu"
	"github.com/ian-howell/tuitris/model/options"
	"github.com/ian-howell/tuitris/model/play"
	"github.com/ian-howell/tuitris/model/splash"
	"github.com/ian-howell/tuitris/model/win"
	"github.com/ian-howell/tuitris/screen"
	"github.com/ian-howell/tuitris/styles"
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
	PlayModel     play.Model
	WinModel      win.Model
	LoseModel     lose.Model
	ErrorModel    iamerror.Model
}

const (
	mainWidth  = 62
	mainHeight = 46
)

func New() (m Model, retErr error) {
	mvp := viewport.New(mainWidth, mainHeight)
	mvp.Style = styles.RoundedPurpleBorder()

	splashModel, err := splash.New()
	retErr = errors.Join(retErr, err)

	mainMenuModel, err := mainmenu.New()
	retErr = errors.Join(retErr, err)

	optionsModel, err := options.New()
	retErr = errors.Join(retErr, err)

	playModel, err := play.New()
	retErr = errors.Join(retErr, err)

	winModel, err := win.New()
	retErr = errors.Join(retErr, err)

	loseModel, err := lose.New()
	retErr = errors.Join(retErr, err)

	errorModel, err := iamerror.New()
	retErr = errors.Join(retErr, err)

	m = Model{
		CurrentScreen: screen.Splash,

		MainViewport: mvp,

		SplashModel:   splashModel,
		MainMenuModel: mainMenuModel,
		OptionsModel:  optionsModel,
		PlayModel:     playModel,
		WinModel:      winModel,
		LoseModel:     loseModel,
		ErrorModel:    errorModel,
	}

	return m, retErr
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
	case screen.Play:
		m.PlayModel, m.CurrentScreen = m.PlayModel.Update(msg)
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
	case screen.Play:
		s = m.PlayModel.View()
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
