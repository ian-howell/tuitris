package main

import (
	"fmt"
	"os"

	"github.com/ian-howell/tuitris/ring"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	splashMenu, err := ring.New(
		Choice{
			Name:       "Menu",
			Cmd:        MenuCmd,
			NextScreen: MenuScreen,
		},
	)
	checkError(err)

	menuMenu, err := ring.New(
		Choice{
			Name:       "Init",
			Cmd:        InitCmd,
			NextScreen: InitScreen,
		},
		Choice{
			Name:       "Options",
			Cmd:        OptionsCmd,
			NextScreen: OptionsScreen,
		},
		Choice{
			Name:       "Exit",
			Cmd:        ExitCmd,
			NextScreen: ErrorScreen,
		},
	)
	checkError(err)

	optionsMenu, err := ring.New(
		Choice{
			Name:       "Menu",
			Cmd:        MenuCmd,
			NextScreen: MenuScreen,
		},
	)
	checkError(err)

	initScreen, err := ring.New(
		Choice{
			Name:       "Play",
			Cmd:        PlayCmd,
			NextScreen: PlayScreen,
		},
	)
	checkError(err)

	playScreen, err := ring.New(
		Choice{
			Name:       "Pause",
			Cmd:        PauseCmd,
			NextScreen: PauseScreen,
		},
		Choice{
			Name:       "Win",
			Cmd:        WinCmd,
			NextScreen: WinScreen,
		},
		Choice{
			Name:       "Lose",
			Cmd:        LoseCmd,
			NextScreen: LoseScreen,
		},
	)
	checkError(err)

	pauseScreen, err := ring.New(
		Choice{
			Name:       "Init",
			Cmd:        InitCmd,
			NextScreen: InitScreen,
		},
		Choice{
			Name:       "Menu",
			Cmd:        MenuCmd,
			NextScreen: MenuScreen,
		},
		Choice{
			Name:       "Exit",
			Cmd:        ExitCmd,
			NextScreen: ErrorScreen,
		},
	)
	checkError(err)

	winScreen, err := ring.New(
		Choice{
			Name:       "Init",
			Cmd:        InitCmd,
			NextScreen: InitScreen,
		},
		Choice{
			Name:       "Menu",
			Cmd:        MenuCmd,
			NextScreen: MenuScreen,
		},
		Choice{
			Name:       "Exit",
			Cmd:        ExitCmd,
			NextScreen: ErrorScreen,
		},
	)
	checkError(err)

	loseScreen, err := ring.New(
		Choice{
			Name:       "Init",
			Cmd:        InitCmd,
			NextScreen: InitScreen,
		},
		Choice{
			Name:       "Menu",
			Cmd:        MenuCmd,
			NextScreen: MenuScreen,
		},
		Choice{
			Name:       "Exit",
			Cmd:        ExitCmd,
			NextScreen: ErrorScreen,
		},
	)
	checkError(err)

	initialModel := Model{
		CurrentScreen: SplashScreen,
		Menus: map[Screen]ring.Ring[Choice]{
			SplashScreen:  splashMenu,
			MenuScreen:    menuMenu,
			OptionsScreen: optionsMenu,
			InitScreen:    initScreen,
			PlayScreen:    playScreen,
			PauseScreen:   pauseScreen,
			WinScreen:     winScreen,
			LoseScreen:    loseScreen,
		},
	}

	p := tea.NewProgram(
		initialModel,
		tea.WithAltScreen(),
	)

	_, err = p.Run()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}

}
