package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

//go:generate stringer -type Screen
type Screen int

const (
	ErrorScreen Screen = iota
	SplashScreen
	MenuScreen
	OptionsScreen
	PlayScreen
	InitScreen
	PauseScreen
	WinScreen
	LoseScreen
)

func (s Screen) HasMenu() bool {
	switch s {
	case SplashScreen,
		MenuScreen,
		PlayScreen,
		OptionsScreen,
		InitScreen,
		PauseScreen,
		WinScreen,
		LoseScreen:
		return true

	}
	return false
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

var _ tea.Msg = PlayCmd

func PlayCmd() tea.Msg {
	return PlayScreen
}

var _ tea.Msg = OptionsCmd

func OptionsCmd() tea.Msg {
	return OptionsScreen
}

var _ tea.Msg = PauseCmd

func PauseCmd() tea.Msg {
	return PauseScreen
}

var _ tea.Msg = WinCmd

func WinCmd() tea.Msg {
	return WinScreen
}

var _ tea.Msg = LoseCmd

func LoseCmd() tea.Msg {
	return LoseScreen
}

var ExitCmd = tea.Quit
