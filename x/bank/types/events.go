package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// bank module event types
const (
	EventTypeTransfer = "transfer"

	AttributeKeyRecipient = "recipient"
	AttributeKeySender    = sdk.AttributeKeySender

	// supply and balance tracking events name and attributes
	EventTypeCoinSpent    = "coin_spent"
	EventTypeCoinReceived = "coin_received"
	EventTypeCoinMint     = "coinbase" // NOTE(fdymylja): using mint clashes with mint module event
	EventTypeCoinBurn     = "burn"

	AttributeKeySpender  = "spender"
	AttributeKeyReceiver = "receiver"
	AttributeKeyMinter   = "minter"
	AttributeKeyBurner   = "burner"

	// EVM
	EventTypeWeiSpent    = "wei_spent"
	EventTypeWeiReceived = "wei_received"
	EventTypeWeiTransfer = "wei_transfer"
)

// NewCoinSpentEvent constructs a new coin spent sdk.Event
func NewCoinSpentEvent(spender sdk.AccAddress, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		EventTypeCoinSpent,
		sdk.NewAttribute(AttributeKeySpender, spender.String()),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	)
}

// NewCoinReceivedEvent constructs a new coin received sdk.Event
func NewCoinReceivedEvent(receiver sdk.AccAddress, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		EventTypeCoinReceived,
		sdk.NewAttribute(AttributeKeyReceiver, receiver.String()),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	)
}

// NewCoinMintEvent construct a new coin minted sdk.Event
func NewCoinMintEvent(minter sdk.AccAddress, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		EventTypeCoinMint,
		sdk.NewAttribute(AttributeKeyMinter, minter.String()),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	)
}

// NewCoinBurnEvent constructs a new coin burned sdk.Event
func NewCoinBurnEvent(burner sdk.AccAddress, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		EventTypeCoinBurn,
		sdk.NewAttribute(AttributeKeyBurner, burner.String()),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	)
}

// NewWeiSpentEvent constructs a new wei spent sdk.Event
func NewWeiSpentEvent(spender sdk.AccAddress, amount math.Int) sdk.Event {
	return sdk.NewEvent(
		EventTypeWeiSpent,
		sdk.NewAttribute(AttributeKeySpender, spender.String()),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	)
}

// NewWeiReceivedEvent constructs a new wei received sdk.Event
func NewWeiReceivedEvent(receiver sdk.AccAddress, amount math.Int) sdk.Event {
	return sdk.NewEvent(
		EventTypeWeiReceived,
		sdk.NewAttribute(AttributeKeyReceiver, receiver.String()),
		sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
	)
}
