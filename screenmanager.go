package main

import tea "github.com/charmbracelet/bubbletea"

type Screen int

const (
	ScreenError = iota
	ScreenSplash
	ScreenMenu
	ScreenOptions
	ScreenPlay
	ScreenInit
	ScreenExit
	ScreenPause
	ScreenWin
	ScreenLose
)

var Screens = map[Screen]struct{}{}

type ScreenMsg struct{}

func (s ScreenMsg) Next(model Model) tea.Cmd {
	switch model.CurrentScreen {
	case ScreenSplash:
		return model.Splash
	}

	return model.Error
}

type ScreenFunc func(Model) Screen

func (m Model) Splash() tea.Msg { return ScreenError }
func (m Model) Error() tea.Msg  { return ScreenError }
