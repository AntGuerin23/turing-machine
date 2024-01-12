package machine

import "errors"

// Head is the main part of the machine, who navigates the machine, reads and write
type Head struct {
	Tape           []string
	cursorPosition int
}

// MoveLeft throws an error if you try to move past the left boundary, unlike MoveRight where the right side is infinite
func (head *Head) MoveLeft() error {
	if head.cursorPosition-1 < 0 {
		return errors.New("Error: Cannot move past the left boundary")
	}
	head.cursorPosition -= 1
	return nil
}

func (head *Head) MoveRight() {
	if head.cursorPosition+1 >= len(head.Tape) {
		head.extendTape(1)
	}
	head.cursorPosition += 1
}

func (head *Head) Read() string {
	return head.Tape[head.cursorPosition]
}

func (head *Head) Write(symbol string) {
	head.Tape[head.cursorPosition] = symbol
}

func (head *Head) extendTape(sizeAdded int) {
	addedItems := make([]string, sizeAdded)
	head.Tape = append(head.Tape, addedItems...)
}
