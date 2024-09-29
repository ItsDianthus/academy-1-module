package domain

// Session - модель основного игрового процесса
type Session struct {
	SessionMode    Mode
	LastTriesCount int
	Word           string
	Category       string
	Difficulty     int
	LettersUsed    map[rune]bool
	Data           []Category
	GameField      string
}

type Mode int

const (
	GetCategory Mode = iota
	GetDifficulty
	MainGame
	End
)
