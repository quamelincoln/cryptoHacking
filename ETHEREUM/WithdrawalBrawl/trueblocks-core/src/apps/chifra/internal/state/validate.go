// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package statePkg

import (
	"errors"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
)

func (opts *StateOptions) ValidateState() error {
	opts.TestLog()

	if opts.BadFlag != nil {
		return opts.BadFlag
	}

	for _, part := range opts.Parts {
		err := validate.ValidateEnum("--parts", part, "[none|some|all|balance|nonce|code|storage|deployed|accttype]")
		if err != nil {
			return err
		}
	}

	if len(opts.Globals.File) > 0 {
		// do nothing for now

	} else {
		if len(opts.Call) > 0 {
			if len(opts.Parts) > 0 {
				return validate.Usage("The {0} option is not available{1}.", "--parts", " with the --call option")
			}

			if opts.Changes {
				return validate.Usage("The {0} option is not available{1}.", "--changes", " with the --call option")
			}

			if opts.NoZero {
				return validate.Usage("The {0} option is not available{1}.", "--no_zero", " with the --call option")
			}

			for _, addr := range opts.Addrs {
				if validate.IsValidAddress(addr) {
					return validate.Usage("The {0} option is not available{1}.", "--call", " when an address is present")
				}
			}

		} else {

			if len(opts.ProxyFor) > 0 {
				return validate.Usage("The --proxy_for option is only available with the --call option.")
			}

			err := validate.ValidateAtLeastOneAddr(opts.Addrs)
			if err != nil {
				return err
			}

		}
	}

	// Blocks are available for both commands
	if len(opts.Blocks) > 0 {
		err := validate.ValidateIdentifiers(
			opts.Globals.Chain,
			opts.Blocks,
			validate.ValidBlockIdWithRange,
			1,
		)

		if err != nil {
			if invalidLiteral, ok := err.(*validate.InvalidIdentifierLiteralError); ok {
				return invalidLiteral
			}

			if errors.Is(err, validate.ErrTooManyRanges) {
				return validate.Usage("Specify only a single block range at a time.")
			}

			return err
		}
	}

	return opts.Globals.ValidateGlobals()
}
