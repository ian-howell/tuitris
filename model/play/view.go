package play

import (
	"fmt"
	"strings"

	"github.com/ian-howell/tuitris/styles"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

const (
	playFieldWidth  = 40
	playFieldHeight = 44

	holdWidth  = 8
	holdHeight = 8

	scoreWidth  = 8
	scoreHeight = 28

	queueWidth  = 8
	queueHeight = 42

	spacer = ""
)

func (m Model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Center,
			spacer,
			m.holdView(),
			"HOLD\n\n\n",
			m.scoreView(),
			"SCORE",
		),
		m.playfieldView(),
		lipgloss.JoinVertical(
			lipgloss.Center,
			spacer,
			m.queueView(),
			"QUEUE",
		),
	)
}

func (m Model) holdView() string {
	vp := viewport.New(holdWidth, holdHeight)
	vp.Style = styles.RoundedPurpleBorder()
	return vp.View()
}

func (m Model) scoreView() string {
	vp := viewport.New(scoreWidth, scoreHeight)
	vp.Style = styles.RoundedPurpleBorder()
	return vp.View()
}

func (m Model) playfieldView() string {
	vp := viewport.New(playFieldWidth, playFieldHeight)
	vp.Style = styles.RoundedPurpleBorder()

	sb := strings.Builder{}
	if m.Paused() {
		sb.WriteString(m.pauseView())
	} else {
		sb.WriteString(m.playView())
	}

	vp.SetContent(sb.String())

	return vp.View()
}

func (m Model) playView() string {
	// Don't show the bottom row, it's effectively a "floor"
	inputRows := m.board[:len(m.board)-1]
	grid := rowView(inputRows[0])
	for _, inputRow := range inputRows[1:] {
		grid = lipgloss.JoinVertical(lipgloss.Center, grid, rowView(inputRow))
	}
	return grid
}

func rowView(inputRow []rune) string {
	var row string
	// Each row has "walls", so let's not print those
	for _, c := range inputRow[1 : len(inputRow)-1] {
		row = lipgloss.JoinHorizontal(lipgloss.Center, row, cellView(c))
	}
	return row
}

func cellView(c rune) string {
	if c == 'W' {
		return "████\n████"
	}
	s := string(c)
	return strings.Join([]string{strings.Repeat(s, 4), strings.Repeat(s, 4)}, "\n")
}

func (m Model) pauseView() string {
	s := "Which screen should we go to next?\n\n"
	for i, choice := range m.PauseMenu.Values() {
		cursor := " "
		if m.PauseMenu.Cursor() == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	return s
}

func (m Model) queueView() string {
	vp := viewport.New(queueWidth, queueHeight)
	vp.Style = styles.RoundedPurpleBorder()
	return vp.View()
}
