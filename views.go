package main

import "fmt"

func ViewMenu(model Model) string {
	s := "Which screen should we go to next?\n\n"

	// Iterate over our choices
	for i, choice := range model.MenuMenu.Choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if model.MenuMenu.Cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice.Name())
	}

	// Send the UI for rendering
	return s
}
