package types

import errorsmod "cosmossdk.io/errors"

// x/nemodist errors
var (
	ErrInvalidProposalAmount  = errorsmod.Register(ModuleName, 2, "invalid community pool multi-spend proposal amount")
	ErrEmptyProposalRecipient = errorsmod.Register(ModuleName, 3, "invalid community pool multi-spend proposal recipient")
)
