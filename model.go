package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ian-howell/tuitris/play"
	"github.com/ian-howell/tuitris/ring"
	"github.com/ian-howell/tuitris/screen"
)

const FPS = 1.0 / time.Second

// The Model contains the game's state.
//
// There should only ever be a single instance of a Model.
type Model struct {
	CurrentScreen screen.Screen

	Menus map[screen.Screen]ring.Ring[screen.Screen]

	MainViewport viewport.Model

	PlayModel play.Model
}

func (m Model) Init() tea.Cmd {
	return doTick()
}

// Update updates the game state. This happens in a goroutine alongside the View function.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TickMsg:
		return m, doTick()
	case tea.KeyMsg:
		key := msg.String()
		if key == "ctrl+c" || key == "q" {
			return m, tea.Quit
		}

		cmd := m.HandleKeyPress(key)
		if cmd != nil {
			return m, cmd
		}
	}

	// Update the view
	var cmd tea.Cmd

	switch {
	case m.CurrentScreen == screen.Play:
	case m.CurrentScreen.HasMenu():
		m.MainViewport.SetContent(m.ViewMenuForCurrentScreen())
	}

	m.MainViewport, cmd = m.MainViewport.Update(msg)
	if cmd != nil {
		return m, cmd
	}

	return m, cmd
}

func (m Model) View() string {
	var s string
	switch m.CurrentScreen {
	case screen.Play:
		s = m.PlayModel.View()
	case screen.Error:
		s = "ERROR"
	default:
		s = m.ViewMenuForCurrentScreen()

	}

	m.MainViewport.SetContent(s)
	return m.MainViewport.View()
}

func (m *Model) HandleKeyPress(key string) tea.Cmd {
	if m.CurrentScreen == screen.Play {
		m.HandlePlayScreen(key)
		return nil
	}

	if m.CurrentScreen.HasMenu() {
		m.HandleMenu(key)
		return nil
	}

	return nil
}

func (m *Model) HandlePlayScreen(key string) {
	m.HandleMenu(key)
}

func (m *Model) HandleMenu(key string) {
	menu, ok := m.Menus[m.CurrentScreen]
	if !ok || !m.CurrentScreen.HasMenu() {
		return
	}

	switch key {
	case "up", "k":
		menu.Prev()
	case "down", "j":
		menu.Next()
	case " ":
		m.CurrentScreen = menu.Get()
	}
}

func (m *Model) ViewMenuForCurrentScreen() string {
	menu, ok := m.Menus[m.CurrentScreen]
	if !ok {
		return "ERROR"
	}

	s := fmt.Sprintf("Current Screen: %s\n", m.CurrentScreen)
	s += "Which screen should we go to next?\n\n"

	// Iterate over our choices
	for i, choice := range menu.Values() {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if menu.Cursor() == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// Send the UI for rendering
	return s
}
