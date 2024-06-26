package play

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ian-howell/tuitris/screen"
)

func (m Model) Update(msg tea.Msg) (Model, screen.Screen) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, screen.Play
	}

	key := keyMsg.String()
	switch key {
	case "p":
		return m, screen.Pause
	}

	return m, screen.Play
}
