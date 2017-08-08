package ticTacToe

import (
	"strconv"

	"../def"
)

type status uint8
type moveVal uint8

const constPlayer1MoveVal moveVal = 3
const constPlayer2MoveVal moveVal = 5
const constNotMovedVal moveVal = 0

const (
	statusOngoing status = iota
	statusDraw
	statusPlayer1Win
	statusPlayer2Win
)

//Board  q
type Board struct {
	cells        [][]moveVal
	size         uint8
	status       status
	numFreeCells uint8
}

//NewBoard  a
func NewBoard(size uint8) *Board {
	board := Board{}

	board.init(size)
	return &board
}

// Reinit init the board to pre-start game status
func (board *Board) ReInit() {
	board.init(board.size)
}

// init the board
func (board *Board) init(size uint8) {
	board.status = statusOngoing

	board.cells = make([][]moveVal, size)
	for i := range board.cells {
		board.cells[i] = make([]moveVal, size)
		for j := range board.cells[i] {
			board.cells[i][j] = constNotMovedVal
		}
	}
	board.size = size
	board.numFreeCells = size * size
}

// PlayerMove Make a move.
func (board *Board) PlayerMove(player def.Player, input uint8) bool {
	row, col := board.getRowAndColFromInput(input)
	if !board.canSelectCell(row, col) {
		return false
	}
	if player == def.Player1 {
		board.cells[row][col] = constPlayer1MoveVal
	} else {
		board.cells[row][col] = constPlayer2MoveVal
	}
	board.numFreeCells = board.numFreeCells - 1
	board.updateStatus(row, col)
	return true
}

// updateStatus updates the board status. Checks if the game is ended
// @var row uint8 The Row that was selected
// @var col uint8 the Col that was selected
func (board *Board) updateStatus(row uint8, col uint8) {

	// only check the horizontal and vertical that the cell selected affects
	board.status = board.doCheckHorizontal(row)
	if board.IsEndGame() {
		return
	}
	board.status = board.doCheckVertical(col)
	if board.IsEndGame() {
		return
	}

	//for diagonal checks, always check, since I have not yet added a check for whether the slot is in one of the affected cells
	// check in this direction: /
	board.status = board.doCheckDiagonal(true)
	if board.IsEndGame() {
		return
	}

	// check in this direction: \
	board.status = board.doCheckDiagonal(false)
	if board.IsEndGame() {
		return
	}
	if board.numFreeCells == 0 {
		board.status = statusDraw
		return
	}
}

// CanAllowPlayerInput check if player input is allowed
func (board *Board) CanAllowPlayerInput(input uint8) bool {

	if input < 1 || input > board.LargestInputAllowed() {
		return false
	}

	row, col := board.getRowAndColFromInput(input)
	return board.canSelectCell(row, col)
}

// canSelectCell given row and column, check if player can select that cell
func (board *Board) canSelectCell(row uint8, col uint8) bool {
	if row >= board.size {
		return false
	}
	if col >= board.size {
		return false
	}
	return board.cells[row][col] == constNotMovedVal
}

func (board *Board) getRowAndColFromInput(input uint8) (uint8, uint8) {
	return (input - 1) / board.size, (input - 1) % board.size
}

// String method to allow printing of Board object
func (board *Board) String() string {
	ret := ""

	for i := range board.cells {
		if i > 0 {
			ret += "\n"
		}
		// board.cells[i] = make([]uint8, numCols)
		for j := range board.cells[i] {
			// board.cells[i][j] = 1
			ret += board.getSlotOutputValue(i, j) + " "
		}
	}
	return ret
}

/***************************************************************************
* check win condition for board.
****************************************************************************/

//check board from left to right
func (board *Board) doCheckHorizontal(row uint8) status {
	firstVal := board.cells[row][0]
	return board.checkHorizontal(row, 1, firstVal)
}

func (board *Board) checkHorizontal(row uint8, curCol uint8, prevVal moveVal) status {
	currVal := board.cells[row][curCol]
	// if cell different from previous cell, that means this row no one wins. just returns
	if currVal != prevVal {
		return statusOngoing
	}
	//last cell. that means someone has got 3 in a row
	if curCol == board.size-1 {
		return board.valueToStatus(currVal)
	}
	return board.checkHorizontal(row, curCol+1, currVal)
}

func (board *Board) doCheckVertical(col uint8) status {
	firstVal := board.cells[0][col]
	return board.checkVertical(1, col, firstVal)
}

func (board *Board) checkVertical(curRow uint8, col uint8, prevVal moveVal) status {

	currVal := board.cells[curRow][col]
	// if cell different from previous cell, that means this row no one wins. just returns
	if currVal != prevVal {
		return statusOngoing
	}
	//last cell. that means someone has got 3 in a row
	if curRow == board.size-1 {
		return board.valueToStatus(currVal)
	}
	return board.checkVertical(curRow+1, col, currVal)
}

// if isReverse is false, check from top left to borrom right
// if isReverse is true, check from bottom left to top right
func (board *Board) doCheckDiagonal(isReverse bool) status {
	var firstVal moveVal
	if !isReverse {
		firstVal = board.cells[0][0]
		return board.checkDiagonal(1, 1, isReverse, firstVal)
	}

	firstVal = board.cells[board.size-1][0]
	return board.checkDiagonal(board.size-2, 1, isReverse, firstVal)
}

func (board *Board) checkDiagonal(row uint8, col uint8, isReverse bool, prevVal moveVal) status {
	currVal := board.cells[row][col]

	// if cell different from previous cell, that means this row no one wins. just returns
	if currVal != prevVal {
		return statusOngoing
	}
	//last cell. that means someone has got 3 in a row
	if col == board.size-1 {
		return board.valueToStatus(currVal)
	}
	if isReverse {
		return board.checkDiagonal(row-1, col+1, isReverse, currVal)
	}

	return board.checkDiagonal(row+1, col+1, isReverse, currVal)
}

func (board *Board) valueToStatus(move moveVal) status {
	switch move {
	case constPlayer1MoveVal:
		return statusPlayer1Win
	case constPlayer2MoveVal:
		return statusPlayer2Win
	default:
		return statusOngoing
	}
}

/****************************************************************************
* getSlotOutputValue Get the output, to show the current status of a particular cell.
* e.g
* [O] == Player 1
* [X] == Player 2
* If no player has played that slot, it shows the value player can input to use that slow
*****************************************************************************/
func (board *Board) getSlotOutputValue(row int, col int) string {
	if board.cells[row][col] == constPlayer1MoveVal {
		return board.getPlayer1Value()
	}
	if board.cells[row][col] == constPlayer2MoveVal {
		return board.getPlayer2Value()
	}
	return "[" + strconv.Itoa(int(board.size)*row+col+1) + "]"
}

// GetPlayerValue Get the output value of the player in the slot
// (is it [X] or [O])
func (board *Board) GetPlayerValue(player def.Player) string {
	if player == def.Player1 {
		return board.getPlayer1Value()
	} else {
		return board.getPlayer2Value()
	}
}

func (board *Board) getPlayer1Value() string {
	return " O "
}

func (board *Board) getPlayer2Value() string {
	return " X "
}

/****************************************************************************
* Endgame status check
*****************************************************************************/

// IsEndGame check if this game has ended
func (board *Board) IsEndGame() bool {
	return board.status != statusOngoing
}

// IsPlayer1Win check if player 1 is the winner
func (board *Board) IsPlayer1Win() bool {
	return board.status == statusPlayer1Win
}

// IsPlayer2Win check if player 2 is the winner
func (board *Board) IsPlayer2Win() bool {
	return board.status == statusPlayer2Win
}

// IsDraw check if the game is a draw
func (board *Board) IsDraw() bool {
	return board.status == statusDraw
}

func (board *Board) LargestInputAllowed() uint8 {
	return board.size * board.size
}
