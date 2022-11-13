package server

type Hand struct {
	cards [3]Card
	cardsNotSelected [3]Card
}

const (
	BASTO = 0
	ORO = 1
	COPA = 2
	ESPADA = 3
)

func (hand Hand*) removeCardSelected(posToDelete){
	hand.cardsNotSelected = append(hand.cardsNotSelected[:posTodelete], hand.cardsNotSelected[posTodelete+1:]...)
}

func (hand *Hand) pointsForSuit() int {
	points := 0
	basto := [3]Card
	oro := [3]card
	copa := [3]Card
	espada := [3]Card
	suits := [4]int{basto,oro,copa,espada} 
	repeatedSuit := nil
	 

	for card, _ : range hand.cards {
		if card.suit == "basto"{
			suits[BASTO] = append(suits[BASTO], card)
		}else if card.suit == "oro"{
			suits[ORO] = append(suits[ORO], card)
		}else if cards.suit == "copa"{
			suits[COPA] = append(suits[COPA], card)
		}else{
			suits[ESPADA] = append(suits[ESPADA], card)
		}
	}

	for suit, _ := range suits {
		if len(suit) >= 2 {
			repeatedSuit = suit
			points += 20
		}
	} 

	if repeatedSuit {
		greatestCardNumber := 0
		secondGreatestCardNumber := 0

		if repeatedSuit[0].value < 10 {
			greatestCardNumber = repeatedSuit[0].value
		}
		if repeatedSuit[0].value < 10 {
			secondGreatestCardNumber = repeatedSuit[1].value
		}

		if len(suit) == 3 and repeatedSuit[2].value < 10{
			if greatestCardNumber > repeatedSuit[2].value && secondGreatestCardNumber < repeatedSuit[2].value {
				secondGreatestCardNumber := repeatedSuit[2].value
			}else if greatestCardNumber < repeatedSuit[2].value && secondGreatestCardNumber > repeatedSuit[2].value {
				greatestCardNumber := repeatedSuit[2].value
			}else if greatestCardNumber < repeatedSuit[2].value && secondGreatestCardNumber < repeatedSuit[2].value{
				if greatestCardNumber > secondGreatestCardNumber {
					secondGreatestCardNumber := repeatedSuit[2].value
				}else {
					greatestCardNumber := repeatedSuit[2].value
				}
			}
		}	
		points += greatestCardNumber + secondGreatestCardNumber	
	}
}

func (hand *Hand) calculateSum() int {
	pointsForSuit := hand.pointsForSuit()
}

func (hand *Hand) winsOver(otherHand *Hand) bool {
	sumForHand := hand.calculateSum()
	sumForOtherHand := otherHand.calculateSum()

	if sumForHand >= sumForOtherHand {
		//si es empate gana el que es mano en la ronda
		return true
	}
	return false
}
