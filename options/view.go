package options

import (
	"fmt"
)

func (m Model) View() string {
	s := "Current Screen: Options\n"
	s += "Which screen should we go to next?\n\n"
	for i, choice := range m.Menu.Values() {
		cursor := " "
		if m.Menu.Cursor() == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	return s
}
