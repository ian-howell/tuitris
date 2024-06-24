package mainmenu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ian-howell/tuitris/screen"
)

func (m Model) Update(msg tea.Msg) (Model, screen.Screen) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, screen.MainMenu
	}

	key := keyMsg.String()
	switch key {
	case "up", "k":
		m.Menu.Prev()
	case "down", "j":
		m.Menu.Next()
	case " ":
		return m, m.Menu.Get()
	}

	return m, screen.MainMenu
}
