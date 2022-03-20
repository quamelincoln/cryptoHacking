// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package validate

import (
	tslibPkg "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/tslib"
)

// Let's define bitmasks to make it easier to validate multiple block identifiers
// (which usually are arguments to `chifra` commands). This way we can specify
// that we want e.g. both block number and a special block as valid arguments
type ValidArgumentType uint16

const (
	ValidArgumentBlockHash ValidArgumentType = 1 << iota
	ValidArgumentBlockNumber
	ValidArgumentDate
	ValidArgumentRange
	ValidArgumentSpecialBlock
	ValidArgumentTransHash
	ValidArgumentTransBlockNumberAndId
	ValidArgumentTransBlockHashAndId
)

const ValidTransId = ValidArgumentTransHash | ValidArgumentTransBlockNumberAndId | ValidArgumentTransBlockHashAndId
const ValidBlockId = ValidArgumentBlockHash | ValidArgumentBlockNumber | ValidArgumentSpecialBlock
const ValidBlockIdWithRange = ValidBlockId | ValidArgumentRange
const ValidBlockIdWithRangeAndDate = ValidBlockIdWithRange | ValidArgumentDate

// ValidateIdentifiers validates multiple identifiers against multiple valid types (specified as bitmasks).
// If any of the identifiers are invalid, it returns error
// If all identifiers are valid, it returns nil
//
// Examples:
//     ValidateIdentifiers(identifiers, ValidArgumentBlockNumber | ValidArgumentDate, 0)
//     ValidateIdentifiers(identifiers, ValidArgumentRange, 1)
//
// This routine can be used for both block identifiers and transaction
func ValidateIdentifiers(chain string, identifiers []string, validTypes ValidArgumentType, maxRanges int) error {
	// A helper function to check if bitmask is set
	isBitmaskSet := func(flag ValidArgumentType) bool {
		return validTypes&flag != 0
	}

	rangesFound := 0

	for _, identifier := range identifiers {
		if isBitmaskSet(ValidArgumentBlockHash) && IsBlockHash(identifier) {
			continue
		}

		if isBitmaskSet(ValidArgumentBlockNumber) && IsBlockNumber(identifier) {
			continue
		}

		if isBitmaskSet(ValidArgumentDate) && IsDateTimeString(identifier) {
			if IsBeforeFirstBlock(chain, identifier) {
				return &InvalidIdentifierLiteralError{
					Value: identifier,
					Msg:   "is before the first block.",
				}
			}
			continue
		}

		if isBitmaskSet(ValidArgumentSpecialBlock) && tslibPkg.IsSpecialBlock(chain, identifier) {
			continue
		}

		if isBitmaskSet(ValidArgumentTransHash) && IsTransHash(identifier) {
			continue
		}

		if isBitmaskSet(ValidArgumentTransBlockNumberAndId) && IsTransBlockNumAndId(identifier) {
			continue
		}

		if isBitmaskSet(ValidArgumentTransBlockHashAndId) && IsTransBlockHashAndId(identifier) {
			continue
		}

		// If we are at this point and we don't want a range, it means that
		// an identifier was invalid
		if !isBitmaskSet(ValidArgumentRange) {
			return &InvalidIdentifierLiteralError{
				Value: identifier,
			}
		}

		_, rangeErr := IsRange(chain, identifier)
		if rangeErr != nil {
			return rangeErr
		}

		rangesFound++
		if rangesFound > maxRanges {
			return ErrTooManyRanges
		}
	}

	return nil
}
