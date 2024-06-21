package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Screen int

const (
	ErrorScreen = iota
	SplashScreen
	MenuScreen
	OptionsScreen
	PlayScreen
	InitScreen
	PauseScreen
	WinScreen
	LoseScreen
)

func (m *Model) HandleKeyPress(key string) tea.Cmd {
	switch m.CurrentScreen {
	case MenuScreen:
		currentMenu := m.Menus[m.CurrentScreen]
		switch key {
		case "up", "k":
			currentMenu.Next()
		case "down", "j":
			currentMenu.Prev()
		case " ":
			m.CurrentScreen = currentMenu.Get().NextScreen
			currentMenu.Reset()
			return currentMenu.Get().Cmd
		}
	}
	return nil
}

var _ tea.Cmd = SplashCmd

func SplashCmd() tea.Msg {
	return SplashScreen
}

var _ tea.Cmd = ErrorCmd

func ErrorCmd() tea.Msg {
	return ErrorScreen
}

var _ tea.Cmd = MenuCmd

func MenuCmd() tea.Msg {
	return MenuScreen
}

var _ tea.Msg = InitCmd

func InitCmd() tea.Msg {
	return InitScreen
}
