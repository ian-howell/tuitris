package screen

//go:generate stringer -type Screen
type Screen int

const (
	Error Screen = iota
	Splash
	MainMenu
	Options
	Reset
	Play
	Pause
	Win
	Lose
	Exit
)
