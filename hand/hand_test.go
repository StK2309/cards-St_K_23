package hand

import "fmt"

func ExampleNewHand() {
	hand := NewHand()
	hand.AddCard("Ace of Spades")
	hand.AddCard("10 of Hearts")
	fmt.Println(hand.GetCards())
	fmt.Println(hand.Size())

	hand.RemoveCard("10 of Hearts")
	fmt.Println(hand.GetCards())
	fmt.Println(hand.Size())

	hand.Clear()
	fmt.Println(hand.GetCards())
	fmt.Println(hand.Size())

	// Output:
	// [Ace of Spades 10 of Hearts]
	// 2
	// [Ace of Spades]
	// 1
	// []
	// 0
}

func ExampleAddCard() {
	hand := NewHand()
	hand.AddCard("King of Diamonds")
	fmt.Println(hand.GetCards())

	// Output: [King of Diamonds]
}

func ExampleRemoveCard() {
	hand := NewHand()
	hand.AddCard("Queen of Clubs")
	hand.AddCard("Jack of Spades")
	fmt.Println(hand.GetCards())

	removed := hand.RemoveCard("Jack of Spades")
	fmt.Println(removed)
	fmt.Println(hand.GetCards())

	notRemoved := hand.RemoveCard("10 of Hearts")
	fmt.Println(notRemoved)
	fmt.Println(hand.GetCards())

	// Output:
	// [Queen of Clubs Jack of Spades]
	// true
	// [Queen of Clubs]
	// false
	// [Queen of Clubs]
}

func ExampleGetCards() {
	hand := NewHand()
	hand.AddCard("9 of Diamonds")
	hand.AddCard("8 of Clubs")
	cards := hand.GetCards()
	fmt.Println(cards)

	// Output: [9 of Diamonds 8 of Clubs]
}

func ExampleSize() {
	hand := NewHand()
	fmt.Println(hand.Size())

	hand.AddCard("7 of Hearts")
	fmt.Println(hand.Size())

	hand.AddCard("6 of Spades")
	fmt.Println(hand.Size())

	hand.RemoveCard("7 of Hearts")
	fmt.Println(hand.Size())

	// Output:
	// 0
	// 1
	// 2
	// 1
}

func ExampleClear() {
	hand := NewHand()
	hand.AddCard("5 of Diamonds")
	hand.AddCard("4 of Clubs")
	fmt.Println(hand.GetCards())

	hand.Clear()
	fmt.Println(hand.GetCards())
	fmt.Println(hand.Size())

	// Output:
	// [5 of Diamonds 4 of Clubs]
	// []
	// 0
}
