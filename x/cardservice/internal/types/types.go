package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Initial Starting Price for a name that was never previously owned
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

// Whois is a struct that contains all the metadata of a name
type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

// Returns a new Whois with the minprice as the price
func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}

// implement fmt.Stringer
func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}


type User struct {
	Alias string         		`json:"value"`
	OwnedCards []uint64			`json:"ownedCards"`
	VoteRights []VoteRight	`json:"voteRight"`
}

func NewUser() User {
	return User{
		Alias: "newPlayer",
		OwnedCards: []uint64{},
		VoteRights: []VoteRight{},
	}
}

type Card struct {
	Owner sdk.AccAddress
	Content []byte
	Status string
	VotePool sdk.Coin
	FairEnoughVotes uint64
	OverpoweredVotes uint64
	UnderpoweredVotes uint64
	InappropriateVotes uint64
	Nerflevel int64
}

func NewCard(owner sdk.AccAddress) Card {
	return Card{
		Owner: owner,
		Content: []byte{},
		Status: "scheme",
		VotePool: sdk.NewInt64Coin("credits", 0),
		FairEnoughVotes: 0,
		OverpoweredVotes: 0,
		UnderpoweredVotes: 0,
		InappropriateVotes: 0,
		Nerflevel: 0,
	}
}

type VoteRight struct {
	cardId string
	expireBlock int64
}