package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(FPS, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
