package def

// Player Def for Player1, Player2
type Player uint8

// Mode Game Mode for the player
type Mode uint8

const (
	Player1 Player = iota
	Player2
)

const (
	ModeMin Mode = iota
	ModePlayerVsPlayer
	ModePlayerVsAI
	ModeAIVsAI
	ModeQuit //quit
	ModeMax  //maximum allowed mode value
)

func GetOtherPlayer(player Player) Player {
	if player == Player1 {
		return Player2
	} else {
		return Player1
	}
}
