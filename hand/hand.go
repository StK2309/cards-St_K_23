package hand

// TODO: Hier Datentyp `Hand` definieren
// und Methoden hinzufügen.
// Tests in separater Datei nicht vergessen.

// Hand represents a collection of playing cards
type Hand struct {
	cards []string
}

// NewHand creates and returns a new empty Hand
func NewHand() *Hand {
	return &Hand{
		cards: []string{},
	}
}

// AddCard adds a card to the hand
func (h *Hand) AddCard(card string) {
	h.cards = append(h.cards, card)
}

// RemoveCard removes the first occurrence of a card from the hand
func (h *Hand) RemoveCard(card string) bool {
	for i, c := range h.cards {
		if c == card {
			h.cards = append(h.cards[:i], h.cards[i+1:]...)
			return true
		}
	}
	return false
}

// GetCards returns a copy of all cards in the hand
func (h *Hand) GetCards() []string {
	return append([]string{}, h.cards...)
}

// Size returns the number of cards in the hand
func (h *Hand) Size() int {
	return len(h.cards)
}

// Clear removes all cards from the hand
func (h *Hand) Clear() {
	h.cards = []string{}
}
