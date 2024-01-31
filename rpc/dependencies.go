// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"

	"github.com/AnomalyFi/hypersdk/chain"
	"github.com/AnomalyFi/hypersdk/crypto/ed25519"
	"github.com/AnomalyFi/nodekit-seq/archiver"
	"github.com/AnomalyFi/nodekit-seq/genesis"
	"github.com/AnomalyFi/nodekit-seq/types"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/trace"
)

type Controller interface {
	Genesis() *genesis.Genesis
	Tracer() trace.Tracer
	GetTransaction(context.Context, ids.ID) (bool, int64, bool, chain.Dimensions, uint64, error)
	GetAssetFromState(context.Context, ids.ID) (bool, []byte, uint8, []byte, uint64, ed25519.PublicKey, bool, error)
	GetBalanceFromState(context.Context, ed25519.PublicKey, ids.ID) (uint64, error)
	GetLoanFromState(context.Context, ids.ID, ids.ID) (uint64, error)
	UnitPrices(ctx context.Context) (chain.Dimensions, error)
	Submit(
		ctx context.Context,
		verifySig bool,
		txs []*chain.Transaction,
	) (errs []error)
	GetBlockFromArchiver(
		ctx context.Context,
		dbBlock *archiver.DBBlock,
	) (*chain.StatefulBlock, *ids.ID, error)
	GetByHeight(
		height uint64,
		end int64,
		reply *types.BlockHeadersResponse,
	) error
	GetByID(
		args *types.GetBlockHeadersIDArgs,
		reply *types.BlockHeadersResponse,
	) error
	GetByStart(
		args *types.GetBlockHeadersByStartArgs,
		reply *types.BlockHeadersResponse,
	) error
}
