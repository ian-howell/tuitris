package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Screen int

const (
	ScreenError = iota
	ScreenSplash
	ScreenMenu
	ScreenOptions
	ScreenPlay
	ScreenInit
	ScreenPause
	ScreenWin
	ScreenLose
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

var _ tea.Cmd = CmdSplash

func CmdSplash() tea.Msg {
	return ScreenSplash
}

var _ tea.Cmd = CmdError

func CmdError() tea.Msg {
	return ScreenError
}

var _ tea.Cmd = CmdMenu

func CmdMenu() tea.Msg {
	return ScreenMenu
}

var _ tea.Msg = CmdInit

func CmdInit() tea.Msg {
	return ScreenInit
}
