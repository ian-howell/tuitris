package main

import "fmt"

func ViewMenu(model Model) string {
	s := "Which screen should we go to next?\n\n"

	currentMenu := model.Menus[model.CurrentScreen]
	// Iterate over our choices
	for i, choice := range currentMenu.Values() {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if currentMenu.Cursor() == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice.Name)
	}

	// Send the UI for rendering
	return s
}
