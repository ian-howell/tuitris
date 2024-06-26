package styles

import "github.com/charmbracelet/lipgloss"

func RoundedPurpleBorder() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62"))
}
