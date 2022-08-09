// Copyright 2019 The klaytn Authors
// This file is part of the klaytn library.
//
// The klaytn library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The klaytn library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the klaytn library. If not, see <http://www.gnu.org/licenses/>.

package sc

import (
	"bytes"
	"math/big"

	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	bridgecontract "github.com/klaytn/klaytn/contracts/bridge"
	"github.com/klaytn/klaytn/crypto/secp256k1"
	"github.com/klaytn/klaytn/datasync/downloader"
	"github.com/klaytn/klaytn/networks/p2p"
	"github.com/klaytn/klaytn/params"
	"github.com/klaytn/klaytn/rlp"
	"github.com/pkg/errors"
)

var ErrRPCDecode = errors.New("failed to decode mainbridge rpc call message")

type MainBridgeHandler struct {
	mainbridge *MainBridge
	// childChainID is the first received chainID from child chain peer.
	childChainIDs map[common.Address]*big.Int
}

func NewMainBridgeHandler(scc *SCConfig, main *MainBridge) (*MainBridgeHandler, error) {
	return &MainBridgeHandler{
		mainbridge:    main,
		childChainIDs: make(map[common.Address]*big.Int),
	}, nil
}

func (mbh *MainBridgeHandler) HandleSubMsg(p BridgePeer, msg p2p.Msg) error {
	logger.Trace("mainbridge handle sub message", "msg.Code", msg.Code)

	// Handle the message depending on its contents
	switch msg.Code {
	case ServiceChainCall:
		if err := mbh.handleCallMsg(p, msg); err != nil {
			return err
		}
		return nil
	case StatusMsg:
		return nil
	case ServiceChainTxsMsg:
		logger.Trace("received ServiceChainTxsMsg")
		// TODO-Klaytn how to check acceptTxs
		// Transactions arrived, make sure we have a valid and fresh chain to handle them
		//if atomic.LoadUint32(&pm.acceptTxs) == 0 {
		//	break
		//}
		if err := mbh.handleServiceChainTxDataMsg(p, msg); err != nil {
			return err
		}
	case ServiceChainParentChainInfoRequestMsg:
		logger.Debug("received ServiceChainParentChainInfoRequestMsg")
		if err := mbh.handleServiceChainParentChainInfoRequestMsg(p, msg); err != nil {
			return err
		}
	case ServiceChainReceiptRequestMsg:
		logger.Debug("received ServiceChainReceiptRequestMsg")
		if err := mbh.handleServiceChainReceiptRequestMsg(p, msg); err != nil {
			return err
		}
	case ServiceChainRequestHandleReceiptMsg:
		logger.Debug("[SC] received ServiceChainRequestHandleReceiptMsg")
		if err := mbh.handleServiceChainRequestHandleReceiptMsg(p, msg); err != nil {
			return err
		}
	default:
		return errResp(ErrInvalidMsgCode, "%v", msg.Code)
	}
	return nil
}

func (mbh *MainBridgeHandler) handleCallMsg(p BridgePeer, msg p2p.Msg) error {
	logger.Trace("mainbridge writes the rpc call message to rpc server", "msg.Size", msg.Size, "msg", msg)
	data := make([]byte, msg.Size)
	err := msg.Decode(&data)
	if err != nil {
		logger.Error("error in mainbridge message handler", "err", err)
		return err
	}

	// Write to RPC server pipe
	_, err = mbh.mainbridge.rpcConn.Write(data)
	if err != nil {
		logger.Error("failed to write to the rpc server pipe", "err", err)
		return err
	}
	return nil
}

// handleServiceChainTxDataMsg handles service chain transactions from child chain.
// It will return an error if given tx is not TxTypeChainDataAnchoring type.
func (mbh *MainBridgeHandler) handleServiceChainTxDataMsg(p BridgePeer, msg p2p.Msg) error {
	// pm.txMsgLock.Lock()
	// Transactions can be processed, parse all of them and deliver to the pool
	var txs []*types.Transaction
	if err := msg.Decode(&txs); err != nil {
		return errResp(ErrDecode, "msg %v: %v", msg, err)
	}

	// Only valid txs should be pushed into the pool.
	// Invalid txs found are sent to child chain for further action.
	invalidTxs := make([]InvalidParentChainTx, 0)
	for _, tx := range txs {
		if tx == nil {
			invalidTxs = append(invalidTxs, InvalidParentChainTx{tx.Hash(), errResp(ErrDecode, "tx is nil").Error()})
			continue
		}
		if err := mbh.mainbridge.txPool.AddRemote(tx); err != nil {
			txHash := tx.Hash()
			logger.Trace("Invalid tx found",
				"txType", tx.Type(), "txNonce", tx.Nonce(), "txHash", txHash.String(), "err", err)
			invalidTxs = append(invalidTxs, InvalidParentChainTx{txHash, err.Error()})
		}
	}
	if len(invalidTxs) > 0 {
		return p.SendServiceChainInvalidTxResponse(invalidTxs)
	}
	return nil
}

// handleServiceChainParentChainInfoRequestMsg handles parent chain info request message from child chain.
// It will send the nonce of the account and its gas price to the child chain peer who requested.
func (mbh *MainBridgeHandler) handleServiceChainParentChainInfoRequestMsg(p BridgePeer, msg p2p.Msg) error {
	var addr common.Address
	if err := msg.Decode(&addr); err != nil {
		return errResp(ErrDecode, "msg %v: %v", msg, err)
	}
	nonce := mbh.mainbridge.txPool.GetPendingNonce(addr)
	chainCfg := mbh.mainbridge.blockchain.Config()
	pcInfo := parentChainInfo{
		nonce,
		mbh.mainbridge.txPool.GasPrice().Uint64(),
		params.KIP71Config{
			LowerBoundBaseFee:         chainCfg.Governance.KIP71.LowerBoundBaseFee,
			UpperBoundBaseFee:         chainCfg.Governance.KIP71.UpperBoundBaseFee,
			GasTarget:                 chainCfg.Governance.KIP71.GasTarget,
			MaxBlockGasUsedForBaseFee: chainCfg.Governance.KIP71.MaxBlockGasUsedForBaseFee,
			BaseFeeDenominator:        chainCfg.Governance.KIP71.BaseFeeDenominator,
		},
		chainCfg.IsMagmaForkEnabled(mbh.mainbridge.blockchain.CurrentBlock().Number()),
	}
	p.SendServiceChainInfoResponse(&pcInfo)
	logger.Info("SendServiceChainInfoResponse", "pBridgeAccoount", addr,
		"nonce", pcInfo.Nonce, "gasPrice", pcInfo.GasPrice,
		"LowerBoundBaseFee", pcInfo.KIP71Config.LowerBoundBaseFee,
		"UpperBoundBaseFee", pcInfo.KIP71Config.UpperBoundBaseFee,
		"GasTarget", pcInfo.KIP71Config.GasTarget,
		"MaxBlockGasUsedForBaseFee", pcInfo.KIP71Config.MaxBlockGasUsedForBaseFee,
		"BaseFeeDenominator", pcInfo.KIP71Config.BaseFeeDenominator,
	)
	return nil
}

// handleServiceChainReceiptRequestMsg handles receipt request message from child chain.
// It will find and send corresponding receipts with given transaction hashes.
func (mbh *MainBridgeHandler) handleServiceChainReceiptRequestMsg(p BridgePeer, msg p2p.Msg) error {
	// Decode the retrieval message
	msgStream := rlp.NewStream(msg.Payload, uint64(msg.Size))
	if _, err := msgStream.List(); err != nil {
		return err
	}
	// Gather state data until the fetch or network limits is reached
	var (
		hash               common.Hash
		receiptsForStorage []*types.ReceiptForStorage
	)
	for len(receiptsForStorage) < downloader.MaxReceiptFetch {
		// Retrieve the hash of the next block
		if err := msgStream.Decode(&hash); err == rlp.EOL {
			break
		} else if err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		// Retrieve the receipt of requested service chain tx, skip if unknown.
		receipt := mbh.mainbridge.blockchain.GetReceiptByTxHash(hash)
		if receipt == nil {
			continue
		}
		receiptsForStorage = append(receiptsForStorage, (*types.ReceiptForStorage)(receipt))
	}
	if len(receiptsForStorage) == 0 {
		return nil
	}
	return p.SendServiceChainReceiptResponse(receiptsForStorage)
}

func (mbh *MainBridgeHandler) requestRefund(reqHandleReceipt RequestHandleReceipt) {
	var gasPrice *big.Int
	mb := mbh.mainbridge
	if mb.blockchain.Config().IsMagmaForkEnabled(mb.blockchain.CurrentHeader().Number) {
		gasPrice = new(big.Int).SetUint64(mb.blockchain.Config().Governance.KIP71.UpperBoundBaseFee)
	} else {
		gasPrice = new(big.Int).SetUint64(mb.blockchain.Config().UnitPrice)
	}
	gasLimit := params.UpperGasLimit / 1000000 // allow gas limit to 999999
	var nonce *big.Int = nil
	auth := bind.MakeTransactOptsWithKeystore(mb.acc.keystore, mb.acc.address, nonce, mb.acc.chainID, gasLimit, gasPrice)

	bridgeAddr := reqHandleReceipt.CounterpartBridgeAddr
	bridge, ok := mb.bridges[bridgeAddr]
	if !ok {
		b, err := bridgecontract.NewBridge(bridgeAddr, mb.localBackend)
		if err != nil {
			logger.Error("[SC][MainBridge] Failed to initialize bridge backend. RequestRefund event cannot be emitted", "addr", bridgeAddr, "err", err)
			return
		}
		bridge = b
		mb.bridges[bridgeAddr] = bridge
	}
	reqRefundTx, err := bridge.RequestRefund(auth, reqHandleReceipt.RequestNonce, reqHandleReceipt.RequestTxHash)
	if err != nil {
		logger.Error("[SC][Refund] Failed to create RequestRefund transaction", "requestTxHash", reqHandleReceipt.RequestTxHash, "err", err)
	} else {
		logger.Trace("[SC][Refund] RequestRefund transaction is created",
			"chain", "parent",
			"refundTxHash", reqRefundTx.Hash().Hex(),
			"requestNonce", reqHandleReceipt.RequestNonce,
		)
	}
	mb.reqRefundNonces[reqHandleReceipt.RequestNonce] = true
}

func (mbh *MainBridgeHandler) handleServiceChainRequestHandleReceiptMsg(p BridgePeer, msg p2p.Msg) error {
	var reqHandleReceipt RequestHandleReceipt
	if err := msg.Decode(&reqHandleReceipt); err != nil {
		logger.Error("[SC][P2P] failed to decode", "err", err)
		return errResp(ErrDecode, "msg %v: %v", msg, err)
	}
	handleTxHash := reqHandleReceipt.HandleTxHash[:]
	if pubkey, err := secp256k1.RecoverPubkey(handleTxHash, reqHandleReceipt.HandleTxSig); err == nil {
		if bytes.Compare(pubkey, mbh.mainbridge.accPubkey) == 0 {
			receipt := mbh.mainbridge.blockchain.GetReceiptByTxHash(reqHandleReceipt.HandleTxHash)
			if receipt != nil && receipt.Status != types.ReceiptStatusSuccessful && !mbh.mainbridge.reqRefundNonces[reqHandleReceipt.RequestNonce] { // executed
				mbh.requestRefund(reqHandleReceipt)
			}
		}
	} else {
		logger.Warn("[SC][Handler] Failed to recover a public key", "err", err, "handleTxHash", reqHandleReceipt.HandleTxHash.Hex())
	}
	return nil
}
