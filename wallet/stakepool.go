// Copyright (c) 2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	"github.com/picfight/pfcd/pfcutil"
	"github.com/picfight/pfcwallet/errors"
	"github.com/picfight/pfcwallet/wallet/udb"
	"github.com/picfight/pfcwallet/wallet/walletdb"
)

// StakePoolUserInfo returns the stake pool user information for a user
// identified by their P2SH voting address.
func (w *Wallet) StakePoolUserInfo(userAddress pfcutil.Address) (*udb.StakePoolUser, error) {
	const op errors.Op = "wallet.StakePoolUserInfo"

	switch userAddress.(type) {
	case *pfcutil.AddressPubKeyHash: // ok
	case *pfcutil.AddressScriptHash: // ok
	default:
		return nil, errors.E(op, errors.Invalid, "address must be P2PKH or P2SH")
	}

	var user *udb.StakePoolUser
	err := walletdb.View(w.db, func(tx walletdb.ReadTx) error {
		stakemgrNs := tx.ReadBucket(wstakemgrNamespaceKey)
		var err error
		user, err = w.StakeMgr.StakePoolUserInfo(stakemgrNs, userAddress)
		return err
	})
	if err != nil {
		return nil, errors.E(op, err)
	}
	return user, nil
}
