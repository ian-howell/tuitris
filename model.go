package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const FPS = 1.0 / time.Second

// The Model contains the game's state.
//
// There should only ever be a single instance of a Model.
type Model struct {
	CurrentScreen Screen

	MenuMenu Menu
}

type Menu struct {
	Choices []Choice
	Cursor  int
}

func (m Menu) CurrentChoice() Choice {
	return m.Choices[m.Cursor]
}

type Choice interface {
	Name() string
	Cmd() tea.Cmd
	NextScreen() Screen
}

type InitChoice struct{}

func (InitChoice) Name() string       { return "Init" }
func (InitChoice) Cmd() tea.Cmd       { return CmdInit }
func (InitChoice) NextScreen() Screen { return ScreenInit }

type ExitChoice struct{}

func (ExitChoice) Name() string       { return "Exit" }
func (ExitChoice) Cmd() tea.Cmd       { return tea.Quit }
func (ExitChoice) NextScreen() Screen { return ScreenError }

func (m Model) Init() tea.Cmd {
	return doTick()
}

func (m Model) KeyHandler(screen Screen) KeyHandler {
	switch screen {
	case ScreenMenu:
		return MenuKeyHandler{}
	}
	return MenuKeyHandler{}
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
		cmd := m.KeyHandler(m.CurrentScreen).
			HandleKeyPress(&m, key)
		if cmd != nil {
			return m, cmd
		}
	}

	return m, nil
}

func (m Model) View() string {
	switch m.CurrentScreen {
	case ScreenSplash:
		return "Splash"
	case ScreenMenu:
		return ViewMenu(m)
	case ScreenError:
		return "Error"
	case ScreenInit:
		return "Init"
	}
	return "Error"
}

type TickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(FPS, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
