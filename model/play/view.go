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

const block = "████\n████"

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
	// TODO: Fix the hardcoded numbers
	// rows contains the playing field plus the 2 rows above (the spawn rows)
	rows := make([]string, 0, 22)
	for r := range 22 {
		// cells contains the contents between the walls of the playing field.
		// For this reason, we'll range between 1 and 10 inclusively.
		cells := make([]string, 0, 10)
		for c := 1; c <= 10; c++ {
			cells = append(cells, m.cellView(r, c))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Center, cells...))
	}
	return lipgloss.JoinVertical(lipgloss.Center, rows...)
}

func (m Model) cellView(row, col int) string {
	return renderCell(m.board[row][col])
}

func renderCell(c rune) string {
	style := lipgloss.NewStyle()
	switch c {
	case 'T':
		style = style.Foreground(lipgloss.Color("#FF00FF"))
		return style.Render(block)
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
