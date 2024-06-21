package main

import (
	"time"

	"github.com/ian-howell/tuitris/ring"

	tea "github.com/charmbracelet/bubbletea"
)

const FPS = 1.0 / time.Second

// The Model contains the game's state.
//
// There should only ever be a single instance of a Model.
type Model struct {
	CurrentScreen Screen

	Menus map[Screen]ring.Ring[Choice]
}

type Choice struct {
	Name       string
	Cmd        tea.Cmd
	NextScreen Screen
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
		cmd := m.HandleKeyPress(key)
		if cmd != nil {
			return m, cmd
		}
	}

	return m, nil
}

func (m Model) View() string {
	switch m.CurrentScreen {
	case SplashScreen:
		return "Splash"
	case MenuScreen:
		return DisplayChoices(m)
	case ErrorScreen:
		return "Error"
	case InitScreen:
		return "Init"
	}
	return "Error"
}
