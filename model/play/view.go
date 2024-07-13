package play

import (
	"fmt"
	"strings"

	"github.com/ian-howell/tuitris/styles"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

const (
	mainWidth  = 44 * 2
	mainHeight = 44

	playFieldWidth  = 20 * 2
	playFieldHeight = 40

	holdWidth  = 8 * 2
	holdHeight = 8

	scoreWidth  = 8 * 2
	scoreHeight = 26

	queueWidth  = 8 * 2
	queueHeight = 40
)

func (m Model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.holdView(),
			"HOLD\n\n\n",
			m.scoreView(),
		),
		m.playfieldView(),
		m.queueView(),
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
	var grid string
	inputRows := m.board[:len(m.board)-1]
	for _, inputRow := range inputRows {
		grid = lipgloss.JoinVertical(lipgloss.Center, grid, rowView(inputRow))
	}
	return grid
}

func rowView(inputRow []rune) string {
	var row string
	for _, c := range inputRow {
		row = lipgloss.JoinHorizontal(lipgloss.Center, row, cellView(c))
	}
	return row
}

func cellView(c rune) string {
	const (
		t = "▛▜"
		b = "▙▟"
	)
	if c == 'W' {
		return lipgloss.JoinVertical(lipgloss.Center, t, b)
	}
	s := string(c)
	return fmt.Sprintf("%v%v\n%v%v", s, s, s, s)
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
	vp := viewport.New(queueWidth, holdHeight)
	vp.Style = styles.RoundedPurpleBorder()
	return vp.View()
}
