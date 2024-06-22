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
	if cmd := m.HandleMenu(key); cmd != nil {
		return cmd
	}
	return nil
}

func (m *Model) HandleMenu(key string) tea.Cmd {
	menu, ok := m.Menus[m.CurrentScreen]
	if !ok {
		return nil
	}

	switch key {
	case "up", "k":
		menu.Next()
	case "down", "j":
		menu.Prev()
	case " ":
		m.CurrentScreen = menu.Get().NextScreen
		menu.Reset()
		return menu.Get().Cmd
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
