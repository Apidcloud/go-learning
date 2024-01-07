package main

func main() {
	cards := newDeck()

	//cards.print()
	/*
		oneHand, twoHand := deal(cards, 8)

		oneHand.print()
		twoHand.print()

		err := oneHand.saveToFile("./onehand.txt")

		if err != nil {
			fmt.Println("Something went wrong saving to file")
		} */

	//readHand := newDeckFromFile("./onehand.txt")
	//readHand.print()

	cards.newShuffle()
	cards.print()
}
