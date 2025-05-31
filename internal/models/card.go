package models

// TODO add number of card we have?
type Card struct {
	Name      string
	ImagePath string // Could be an online link or local path.
	// Count represents how many instance of this card we possess.
	Count int
	// Set describes in which set the card is printed.
	Set string
	// SetNumber represents which specific card from the set is represented.
	SetNumber string // Can't be just an integer, because sometime there are things such as "123s"
}

// TODO function that returns a list of options from the non-Zero fields?
