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

	nextScreen := screen.Play
	key := keyMsg.String()
	switch key {
	case "up", "k":
		m.PauseMenu.Prev()
	case "down", "j":
		m.PauseMenu.Next()
	case "p":
		m.Unpause()
	case " ":
		nextScreen = m.PauseMenu.Get()
		m.PauseMenu.Reset()
	}

	return m, nextScreen
}
