package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./def"
	"./player"
	"./ticTacToe"
)

func main() {
	consolereader := bufio.NewReader(os.Stdin)

	board := ticTacToe.NewBoard(3)
	for {
		mode := showMenu(consolereader)
		if mode == def.ModeQuit {
			break
		}
		player1, player2 := generatePlayers(mode, board)

		//game should never come here, if it comes here, there is a problem
		if player1 == nil || player2 == nil {
			fmt.Println("invalid players, please choose mode again..")
			continue
		}

		for {
			playGame(board, consolereader, player1, player2)
			if checkPlayAgain(consolereader) {
				//if player wants to play again, re-initialize the board, wiping it of its previous status
				board.ReInit()
				continue
			}
			break
		}
		board.ReInit()
	}

	fmt.Print("Thank you for playing!")
}

func generatePlayers(mode def.Mode, board *ticTacToe.Board) (player.Player, player.Player) {
	var player1 player.Player
	var player2 player.Player
	switch mode {
	case def.ModePlayerVsPlayer:
		player1 = player.NewHumanPlayer("Player 1", def.Player1, board)
		player2 = player.NewHumanPlayer("Player 2", def.Player2, board)
	case def.ModePlayerVsAI:
		player1 = player.NewHumanPlayer("Player", def.Player1, board)
		player2 = player.NewAIPlayer("AI", def.Player2, board)
	case def.ModeAIVsAI:
		player1 = player.NewAIPlayer("AI1", def.Player1, board)
		player2 = player.NewAIPlayer("AI2", def.Player2, board)
	default:
		player1 = nil
		player2 = nil
	}
	return player1, player2
}

func playGame(board *ticTacToe.Board, consolereader *bufio.Reader, player1 player.Player, player2 player.Player) {

	currentPlayer := player1

	for !board.IsEndGame() {
		//read the input from player.
		// while the input is invalid, keep reading input and attempting to put that move
		for {
			input := currentPlayer.GetNextMove()
			if !board.PlayerMove(currentPlayer.GetPlayerID(), input) {
				fmt.Print("This slot is already taken..: ")
				continue
			}
			break
		}

		//switch player
		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}
	}
	fmt.Println(board)
	if board.IsPlayer1Win() {
		fmt.Println("Player 1 wins!")
	} else if board.IsPlayer2Win() {
		fmt.Println("Player 2 wins!")
	} else if board.IsDraw() {
		fmt.Println("Draw..")
	}
}

func showMenu(consolereader *bufio.Reader) def.Mode {

	fmt.Println(def.ModePlayerVsPlayer, "- Player vs Player")
	fmt.Println(def.ModePlayerVsAI, "- Player vs AI")
	fmt.Println(def.ModeAIVsAI, "- AI vs AI")
	fmt.Println(def.ModeQuit, "- Quit Game")
	fmt.Print("Please choose an option: ")

	for {
		inputString, err := consolereader.ReadString('\n')

		if err != nil {
			fmt.Print("Please enter a correct value1: ")
			continue
		}

		inputString = strings.TrimSuffix(inputString, "\r\n")
		inputInt, err := strconv.ParseInt(inputString, 10, 8)
		if err != nil {
			fmt.Print("Please enter a correct value2: ")
			continue
		}

		inputUint8 := uint8(inputInt)
		inputMode := def.Mode(inputUint8)
		if inputMode <= def.ModeMin || inputMode > def.ModeMax {
			fmt.Print("Invalid option, please input again: ")
			continue
		}
		return inputMode
	}
}

// Check if player wants to play again.
// Returns true if player wants to play again
func checkPlayAgain(consolereader *bufio.Reader) bool {
	for {
		fmt.Print("Play again? (Y/N): ")
		input, err := consolereader.ReadString('\n')

		if err != nil {
			fmt.Print("Please enter a correct value: ")
			continue
		}

		input = strings.TrimSuffix(input, "\r\n")
		if input == "Y" || input == "y" {
			return true
		}
		if input == "N" || input == "n" {
			return false
		}
	}
}
