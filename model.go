package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const FPS = 1.0 / time.Second

// The Model contains the game's state.
//
// There should only ever be a single instance of a Model.
type Model struct {
	CurrentScreen Screen
}

func (m Model) Init() tea.Cmd {
	return doTick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ScreenMsg:
		cmd := msg.Next(m)
		return m, cmd
	case TickMsg:
		return m, doTick()
	}

	return m, nil
}

func (m Model) View() string {
	switch m.CurrentScreen {
	case ScreenSplash:
		return "Splash"
	case ScreenError:
		return "Error"
	}
	return "Error"
}

type TickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(FPS, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
