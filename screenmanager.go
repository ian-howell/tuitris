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

type MenuScreenHandler struct{}

func (MenuScreenHandler) KeyHandler(screen Screen) KeyHandler {
	return MenuKeyHandler{}
}

type KeyHandler interface {
	HandleKeyPress(model *Model, key string) tea.Cmd
}

type MenuKeyHandler struct{}

func (MenuKeyHandler) HandleKeyPress(model *Model, key string) tea.Cmd {
	switch key {
	case "up", "k":
		if model.MenuMenu.Cursor > 0 {
			model.MenuMenu.Cursor--
		}
	case "down", "j":
		if model.MenuMenu.Cursor < len(model.MenuMenu.Choices)-1 {
			model.MenuMenu.Cursor++
		}
	case " ":
		model.CurrentScreen = model.MenuMenu.CurrentChoice().NextScreen()
		return model.MenuMenu.CurrentChoice().Cmd()
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
