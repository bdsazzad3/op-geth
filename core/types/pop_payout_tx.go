// Copyright 2021 The go-ethereum Authors
// Copyright 2023 Bloq, Inc.
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

const (
	PopPayoutTxType = 0x7D

	PoPPayoutSender = "0x8888888888888888888888888888888888888888"
)

var (
	PoPPayoutSenderAddress = common.HexToAddress(PoPPayoutSender)
)

type PopPayoutTx struct {
	// Will always be GovernanceToken contract address
	To *common.Address `rlp:"nil"`

	// gas limit
	Gas uint64

	// ABI-encoded smart contract call to GovernanceToken.mintPoPRewards()
	Data []byte
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *PopPayoutTx) copy() TxData {
	cpy := &PopPayoutTx{
		To:   copyAddressPtr(tx.To),
		Gas:  tx.Gas,
		Data: common.CopyBytes(tx.Data),
	}

	return cpy
}

func (tx *PopPayoutTx) txType() byte           { return PopPayoutTxType }
func (tx *PopPayoutTx) chainID() *big.Int      { return common.Big0 } // Compatibility - Unused
func (tx *PopPayoutTx) accessList() AccessList { return nil }         // Compatibility - Unused
func (tx *PopPayoutTx) data() []byte           { return tx.Data }
func (tx *PopPayoutTx) gas() uint64            { return tx.Gas }
func (tx *PopPayoutTx) gasFeeCap() *big.Int    { return new(big.Int) }
func (tx *PopPayoutTx) gasTipCap() *big.Int    { return new(big.Int) }
func (tx *PopPayoutTx) gasPrice() *big.Int     { return new(big.Int) }
func (tx *PopPayoutTx) value() *big.Int        { return new(big.Int) } // Compatibility - Unused
func (tx *PopPayoutTx) nonce() uint64          { return 0 }            // Compatibility - actual nonce set during execution
func (tx *PopPayoutTx) to() *common.Address    { return tx.To }
func (tx *PopPayoutTx) isSystemTx() bool       { return false } // Compatibility - Unused

func (tx *PopPayoutTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	return dst.Set(new(big.Int))
}

func (tx *PopPayoutTx) rawSignatureValues() (v, r, s *big.Int) {
	return common.Big0, common.Big0, common.Big0
}

func (tx *PopPayoutTx) setSignatureValues(chainID, v, r, s *big.Int) {
	// this is a noop for pop transactions
}

func (tx *PopPayoutTx) encode(b *bytes.Buffer) error {
	return rlp.Encode(b, tx)
}

func (tx *PopPayoutTx) decode(input []byte) error {
	return rlp.DecodeBytes(input, tx)
}
