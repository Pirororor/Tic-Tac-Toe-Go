package player

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"../def"
	"../ticTacToe"
)

type Player interface {
	GetNextMove() uint8
	GetPlayerID() def.Player
	GetName() string
}

// Human player is a playable player! DUH!
type HumanPlayer struct {
	name          string
	playerID      def.Player
	board         *ticTacToe.Board
	consolereader *bufio.Reader
}

// AIPlayer is a computer player, player does not get to control this computer.
type AIPlayer struct {
	name     string
	playerID def.Player
	board    *ticTacToe.Board
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Human player functions
////////////////////////////////////////////////////////////////////////////////////////////////////

// NewHumanPlayer create a new human player
func NewHumanPlayer(name string, playerID def.Player, board *ticTacToe.Board) Player {
	consolereader := bufio.NewReader(os.Stdin)
	player := &HumanPlayer{name: name, playerID: playerID, board: board, consolereader: consolereader}
	return player
}

//GetNextMove Get the next move from the human player
func (player *HumanPlayer) GetNextMove() uint8 {
	token := player.board.GetPlayerValue(player.playerID)
	fmt.Println(player.name + "'s (" + token + ") turn!: ")
	fmt.Println(player.board)
	fmt.Println()
	fmt.Print("Select a place to put : ")
	for {
		input := readPlayerInput(player.consolereader)
		if !player.board.CanAllowPlayerInput(input) {
			fmt.Print("Invalid slot..: ")
			continue
		}
		return input
	}
}

func (player *HumanPlayer) GetPlayerID() def.Player {
	return player.playerID
}

func (player *HumanPlayer) GetName() string {
	return player.name
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// AI player functions
////////////////////////////////////////////////////////////////////////////////////////////////////

// NewAIPlayer create a new AI player
func NewAIPlayer(name string, playerID def.Player, board *ticTacToe.Board) Player {
	player := &AIPlayer{name: name, playerID: playerID, board: board}
	return player
}

//GetNextMove Get the next move from the AI player
func (player *AIPlayer) GetNextMove() uint8 {
	fmt.Println(player.board)
	fmt.Println(player.name + " is thinking...")
	time.Sleep(1 * time.Second)

	for {
		randdom := rand.Intn(int(player.board.LargestInputAllowed()))
		input := uint8(randdom + 1)

		if !player.board.CanAllowPlayerInput(input) {
			continue
		}
		fmt.Println(player.name + " selects " + strconv.Itoa(int(input)))
		return input
	}
}

func (player *AIPlayer) GetPlayerID() def.Player {
	return player.playerID
}

func (player *AIPlayer) GetName() string {
	return player.name
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Common Functions between the AI player and the Human player
////////////////////////////////////////////////////////////////////////////////////////////////////

// Get player input for the board.
// note that it does not check whether the board move is valid or not. (e.g is that space already occupied)
func readPlayerInput(consolereader *bufio.Reader) uint8 {
	for {
		input, err := consolereader.ReadString('\n')

		if err != nil {
			fmt.Print("Please enter a correct value: ")
			continue
		}

		input = strings.TrimSuffix(input, "\r\n")
		inputInt, err := strconv.ParseInt(input, 10, 8)

		//not an integer
		if err != nil {
			fmt.Print("Please enter a correct value: ")
			continue
		}

		inputUint8 := uint8(inputInt)
		return inputUint8
	}
}
