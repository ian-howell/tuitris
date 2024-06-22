package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PlayScreenViewport struct {
	MainViewport      viewport.Model
	HoldViewport      viewport.Model
	ScoreViewport     viewport.Model
	PlayFieldViewport viewport.Model
	QueueViewport     viewport.Model
}

const (
	mainWidth  = 44
	mainHeight = 44

	playFieldWidth  = 20
	playFieldHeight = 40

	holdWidth  = 8
	holdHeight = 8

	scoreWidth  = 8
	scoreHeight = 26

	queueWidth  = 8
	queueHeight = 40
)

func NewPlayScreenViewport() PlayScreenViewport {
	pvp := PlayScreenViewport{
		MainViewport:      viewport.New(mainWidth, mainHeight),
		PlayFieldViewport: viewport.New(playFieldWidth, playFieldHeight),
		HoldViewport:      viewport.New(holdWidth, holdHeight),
		QueueViewport:     viewport.New(queueWidth, queueHeight),
		ScoreViewport:     viewport.New(scoreWidth, scoreHeight),
	}

	pvp.PlayFieldViewport.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62"))

	pvp.HoldViewport.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62"))

	pvp.QueueViewport.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62"))

	pvp.ScoreViewport.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62"))

	return pvp
}

func (p *PlayScreenViewport) Update(model *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	p.PlayFieldViewport.SetContent(model.ViewMenuForCurrentScreen())
	p.PlayFieldViewport, cmd = p.PlayFieldViewport.Update(msg)
	if cmd != nil {
		return model, cmd
	}

	return model, nil
}

func (p PlayScreenViewport) View() string {
	s := lipgloss.JoinHorizontal(
		lipgloss.Top,
		p.HoldViewport.View(),
		p.PlayFieldViewport.View(),
		p.QueueViewport.View(),
	)

	p.MainViewport.SetContent(s)
	return p.MainViewport.View()
}
