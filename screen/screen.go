package screen

//go:generate stringer -type Screen
type Screen int

const (
	Error Screen = iota
	Splash
	MainMenu
	Options
	Play
	Init
	Pause
	Win
	Lose
	Exit
)

func (s Screen) HasMenu() bool {
	switch s {
	case Splash,
		MainMenu,
		Play,
		Options,
		Init,
		Pause,
		Win,
		Lose:
		return true

	}
	return false
}
