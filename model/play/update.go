package play

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ian-howell/tuitris/screen"
)

func (m Model) Update(msg tea.Msg) (Model, screen.Screen) {
	if m.Paused() {
		return m.pauseUpdate(msg)
	}

	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, screen.Play
	}

	key := keyMsg.String()
	switch key {
	case "p":
		m.Pause()
		return m, screen.Play
	}

	return m, screen.Play
}

func (m Model) pauseUpdate(msg tea.Msg) (Model, screen.Screen) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, screen.Win
	}

	key := keyMsg.String()

	if key == "p" {
		m.Unpause()
		return m, screen.Play
	}

	var nextScreen screen.Screen
	switch key {
	case "up", "k":
		m.PauseMenu.Prev()
		return m, screen.Play
	case "down", "j":
		m.PauseMenu.Next()
		return m, screen.Play
	case " ":
		nextScreen = m.PauseMenu.Get()
	default:
		// Ignore all other keys
		return m, screen.Play
	}

	switch nextScreen {
	case screen.Play, screen.MainMenu:
		var err error
		m, err = New()
		if err != nil {
			return m, screen.Error
		}
	}

	return m, nextScreen
}
