// Copyright (c) 2015-2016 The btcsuite developers
// Copyright (c) 2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package cfgutil

import "github.com/picfight/pfcd/pfcutil"

// AddressFlag contains a pfcutil.Address and implements the flags.Marshaler and
// Unmarshaler interfaces so it can be used as a config struct field.
type AddressFlag struct {
	Address pfcutil.Address
}

// NewAddressFlag creates an AddressFlag with a default pfcutil.Address.
func NewAddressFlag(defaultValue pfcutil.Address) *AddressFlag {
	return &AddressFlag{defaultValue}
}

// MarshalFlag satisifes the flags.Marshaler interface.
func (a *AddressFlag) MarshalFlag() (string, error) {
	if a.Address != nil {
		return a.Address.String(), nil
	}

	return "", nil
}

// UnmarshalFlag satisifes the flags.Unmarshaler interface.
func (a *AddressFlag) UnmarshalFlag(addr string) error {
	if addr == "" {
		a.Address = nil
		return nil
	}
	address, err := pfcutil.DecodeAddress(addr)
	if err != nil {
		return err
	}
	a.Address = address
	return nil
}