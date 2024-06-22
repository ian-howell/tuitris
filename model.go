package main

import (
	"fmt"
	"time"

	"github.com/ian-howell/tuitris/ring"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const FPS = 1.0 / time.Second

// The Model contains the game's state.
//
// There should only ever be a single instance of a Model.
type Model struct {
	CurrentScreen Screen

	Menus map[Screen]ring.Ring[Choice]

	PlayViewport viewport.Model
	MainViewport viewport.Model
}

type Choice struct {
	Name       string
	Cmd        tea.Cmd
	NextScreen Screen
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
			return m, ExitCmd
		}

		cmd := m.HandleKeyPress(key)
		if cmd != nil {
			return m, cmd
		}
	}

	// Update the view
	var cmd tea.Cmd

	m.PlayViewport.SetContent(m.ViewMenuForCurrentScreen())
	m.PlayViewport, cmd = m.PlayViewport.Update(msg)
	if cmd != nil {
		return m, cmd
	}

	m.MainViewport.SetContent(m.ViewMenuForCurrentScreen())
	m.MainViewport, cmd = m.MainViewport.Update(msg)
	if cmd != nil {
		return m, cmd
	}

	return m, cmd
}

func (m Model) View() string {
	switch m.CurrentScreen {
	case ErrorScreen:
		return "ERROR"
	case PlayScreen:
		return ViewPlayScreen(&m)

	}
	return m.MainViewport.View()
}

func (m *Model) HandleKeyPress(key string) tea.Cmd {
	if m.CurrentScreen == PlayScreen {
		if cmd := m.HandlePlayScreen(key); cmd != nil {
			return cmd
		}
		return nil
	}

	if m.CurrentScreen.HasMenu() {
		if cmd := m.HandleMenu(key); cmd != nil {
			return cmd
		}
	}
	return nil
}

func (m *Model) HandlePlayScreen(key string) tea.Cmd {
	return m.HandleMenu(key)
}

func (m *Model) HandleMenu(key string) tea.Cmd {
	menu, ok := m.Menus[m.CurrentScreen]
	if !ok || !m.CurrentScreen.HasMenu() {
		return nil
	}

	switch key {
	case "up", "k":
		menu.Prev()
	case "down", "j":
		menu.Next()
	case " ":
		m.CurrentScreen = menu.Get().NextScreen
		cmd := menu.Get().Cmd
		menu.Reset()
		return cmd
	}
	return nil
}

func (m *Model) ViewMenuForCurrentScreen() string {
	menu, ok := m.Menus[m.CurrentScreen]
	if !ok {
		return "ERROR"
	}

	s := "Which screen should we go to next?\n\n"

	// Iterate over our choices
	for i, choice := range menu.Values() {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if menu.Cursor() == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice.Name)
	}

	// Send the UI for rendering
	return s
}
