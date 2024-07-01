package screen

//go:generate stringer -type Screen
type Screen int

const (
	Error Screen = iota
	Splash
	MainMenu
	Options
	Play
	Win
	Lose
	Exit
)
