package reset

import (
	"github.com/ian-howell/tuitris/ring"
	"github.com/ian-howell/tuitris/screen"
)

type Model struct {
	Menu ring.Ring[screen.Screen]
}

func New() (Model, error) {
	menu, err := ring.New(screen.Play)
	if err != nil {
		return Model{}, err
	}
	return Model{Menu: menu}, nil
}
