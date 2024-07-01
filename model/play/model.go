package play

import (
	"github.com/ian-howell/tuitris/ring"
	"github.com/ian-howell/tuitris/screen"
)

type Model struct {
	PauseMenu ring.Ring[screen.Screen]
	isPaused  bool
}

func New() (Model, error) {
	pauseMenu, err := ring.New(screen.Reset, screen.MainMenu, screen.Exit)
	if err != nil {
		return Model{}, err
	}

	return Model{
		PauseMenu: pauseMenu,
	}, nil
}

func (m *Model) Pause() {
	m.isPaused = true
}

func (m *Model) Unpause() {
	m.isPaused = false
}

func (m *Model) Paused() bool {
	return m.isPaused
}
