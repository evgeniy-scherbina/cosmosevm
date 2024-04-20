package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/evgeniy-scherbina/cosmosevm/x/evm/types"
)

// ApplyTransaction runs and attempts to perform a state transition with the given transaction (i.e Message), that will
// only be persisted (committed) to the underlying KVStore if the transaction does not fail.
//
// # Gas tracking
//
// Ethereum consumes gas according to the EVM opcodes instead of general reads and writes to store. Because of this, the
// state transition needs to ignore the SDK gas consumption mechanism defined by the GasKVStore and instead consume the
// amount of gas used by the VM execution. The amount of gas used is tracked by the EVM and returned in the execution
// result.
//
// Prior to the execution, the starting tx gas meter is saved and replaced with an infinite gas meter in a new context
// in order to ignore the SDK gas consumption config values (read, write, has, delete).
// After the execution, the gas used from the message execution will be added to the starting gas consumed, taking into
// consideration the amount of gas returned. Finally, the context is updated with the EVM gas consumed value prior to
// returning.
//
// For relevant discussion see: https://github.com/cosmos/cosmos-sdk/discussions/9072
func (k *Keeper) ApplyTransaction(ctx sdk.Context, tx *ethtypes.Transaction) (*types.MsgEthereumTxResponse, error) {
	//var (
	//	bloom        *big.Int
	//	bloomReceipt ethtypes.Bloom
	//)
	//
	//cfg, err := k.EVMConfig(ctx, sdk.ConsAddress(ctx.BlockHeader().ProposerAddress), k.eip155ChainID)
	//if err != nil {
	//	return nil, errorsmod.Wrap(err, "failed to load evm config")
	//}
	//txConfig := k.TxConfig(ctx, tx.Hash())
	//
	//// get the signer according to the chain rules from the config and block height
	//signer := ethtypes.MakeSigner(cfg.ChainConfig, big.NewInt(ctx.BlockHeight()))
	//msg, err := tx.AsMessage(signer, cfg.BaseFee)
	//if err != nil {
	//	return nil, errorsmod.Wrap(err, "failed to return ethereum transaction as core message")
	//}
	//
	//// snapshot to contain the tx processing and post processing in same scope
	//var commit func()
	//tmpCtx := ctx
	//if k.hooks != nil {
	//	// Create a cache context to revert state when tx hooks fails,
	//	// the cache context is only committed when both tx and hooks executed successfully.
	//	// Didn't use `Snapshot` because the context stack has exponential complexity on certain operations,
	//	// thus restricted to be used only inside `ApplyMessage`.
	//	tmpCtx, commit = ctx.CacheContext()
	//}
	//
	//// pass true to commit the StateDB
	//res, err := k.ApplyMessageWithConfig(tmpCtx, msg, nil, true, cfg, txConfig)
	//if err != nil {
	//	return nil, errorsmod.Wrap(err, "failed to apply ethereum core message")
	//}
	//
	//logs := types.LogsToEthereum(res.Logs)
	//
	//// Compute block bloom filter
	//if len(logs) > 0 {
	//	bloom = k.GetBlockBloomTransient(ctx)
	//	bloom.Or(bloom, big.NewInt(0).SetBytes(ethtypes.LogsBloom(logs)))
	//	bloomReceipt = ethtypes.BytesToBloom(bloom.Bytes())
	//}
	//
	//cumulativeGasUsed := res.GasUsed
	//if ctx.BlockGasMeter() != nil {
	//	limit := ctx.BlockGasMeter().Limit()
	//	cumulativeGasUsed += ctx.BlockGasMeter().GasConsumed()
	//	if cumulativeGasUsed > limit {
	//		cumulativeGasUsed = limit
	//	}
	//}
	//
	//var contractAddr common.Address
	//if msg.To() == nil {
	//	contractAddr = crypto.CreateAddress(msg.From(), msg.Nonce())
	//}
	//
	//receipt := &ethtypes.Receipt{
	//	Type:              tx.Type(),
	//	PostState:         nil, // TODO: intermediate state root
	//	CumulativeGasUsed: cumulativeGasUsed,
	//	Bloom:             bloomReceipt,
	//	Logs:              logs,
	//	TxHash:            txConfig.TxHash,
	//	ContractAddress:   contractAddr,
	//	GasUsed:           res.GasUsed,
	//	BlockHash:         txConfig.BlockHash,
	//	BlockNumber:       big.NewInt(ctx.BlockHeight()),
	//	TransactionIndex:  txConfig.TxIndex,
	//}
	//
	//if !res.Failed() {
	//	receipt.Status = ethtypes.ReceiptStatusSuccessful
	//	// Only call hooks if tx executed successfully.
	//	if err = k.PostTxProcessing(tmpCtx, msg, receipt); err != nil {
	//		// If hooks return error, revert the whole tx.
	//		res.VmError = types.ErrPostTxProcessing.Error()
	//		k.Logger(ctx).Error("tx post processing failed", "error", err)
	//
	//		// If the tx failed in post processing hooks, we should clear the logs
	//		res.Logs = nil
	//	} else if commit != nil {
	//		// PostTxProcessing is successful, commit the tmpCtx
	//		commit()
	//		// Since the post-processing can alter the log, we need to update the result
	//		res.Logs = types.NewLogsFromEth(receipt.Logs)
	//		ctx.EventManager().EmitEvents(tmpCtx.EventManager().Events())
	//	}
	//}
	//
	//// refund gas in order to match the Ethereum gas consumption instead of the default SDK one.
	//if err = k.RefundGas(ctx, msg, msg.Gas()-res.GasUsed, cfg.Params.EvmDenom); err != nil {
	//	return nil, errorsmod.Wrapf(err, "failed to refund gas leftover gas to sender %s", msg.From())
	//}
	//
	//if len(receipt.Logs) > 0 {
	//	// Update transient block bloom filter
	//	k.SetBlockBloomTransient(ctx, receipt.Bloom.Big())
	//	k.SetLogSizeTransient(ctx, uint64(txConfig.LogIndex)+uint64(len(receipt.Logs)))
	//}
	//
	//k.SetTxIndexTransient(ctx, uint64(txConfig.TxIndex)+1)
	//
	//totalGasUsed, err := k.AddTransientGasUsed(ctx, res.GasUsed)
	//if err != nil {
	//	return nil, errorsmod.Wrap(err, "failed to add transient gas used")
	//}
	//
	//// reset the gas meter for current cosmos transaction
	//k.ResetGasMeterAndConsumeGas(ctx, totalGasUsed)
	//return res, nil
}
