package play

import (
	"github.com/ian-howell/tuitris/ring"
	"github.com/ian-howell/tuitris/screen"
)

type Model struct {
	Menu ring.Ring[screen.Screen]
}

func New() (Model, error) {
	menu, err := ring.New(screen.Pause, screen.Win, screen.Lose)
	if err != nil {
		return Model{}, err
	}
	return Model{Menu: menu}, nil
}
