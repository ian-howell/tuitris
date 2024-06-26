package play

import (
	"github.com/ian-howell/tuitris/styles"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

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

	s := "Current Screen: Play\n"
	s += "Press P to pause"

	vp.SetContent(s)

	return vp.View()
}

func (m Model) queueView() string {
	vp := viewport.New(queueWidth, holdHeight)
	vp.Style = styles.RoundedPurpleBorder()
	return vp.View()
}
