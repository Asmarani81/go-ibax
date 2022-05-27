/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package daemons

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/IBAX-io/go-ibax/packages/block"
	"github.com/IBAX-io/go-ibax/packages/common/crypto"
	"github.com/IBAX-io/go-ibax/packages/conf"
	"github.com/IBAX-io/go-ibax/packages/conf/syspar"
	"github.com/IBAX-io/go-ibax/packages/consts"
	"github.com/IBAX-io/go-ibax/packages/protocols"
	"github.com/IBAX-io/go-ibax/packages/service/node"
	"github.com/IBAX-io/go-ibax/packages/storage/sqldb"
	"github.com/IBAX-io/go-ibax/packages/types"
	"github.com/IBAX-io/go-ibax/packages/utils"
	log "github.com/sirupsen/logrus"
	"math"
	"strconv"
	"time"
)

func BlockGeneratorNew(ctx context.Context, d *daemon) error {
	d.sleepTime = time.Second
	if node.IsNodePaused() {
		return nil
	}
	prevBlock := &sqldb.InfoBlock{}
	_, err := prevBlock.Get()
	if err != nil {
		d.logger.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("getting previous block")
		return err
	}
	NodePrivateKey, NodePublicKey := utils.GetNodeKeys()
	if len(NodePrivateKey) < 1 {
		d.logger.WithFields(log.Fields{"type": consts.EmptyObject}).Error("node private key is empty")
		return errors.New(`node private key is empty`)
	}
	if len(NodePublicKey) < 1 {
		d.logger.WithFields(log.Fields{"type": consts.EmptyObject}).Error("node public key is empty")
		return errors.New(`node public key is empty`)
	}

	candidateNode := &sqldb.CandidateNode{}
	var (
		candidateNodes []sqldb.CandidateNode
	)
	candidateNodes, err = candidateNode.GetCandidateNode()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("getting candidate node list")
		return err
	}
	currentCandidateNode, nodePosition := GetThisNodePosition(candidateNodes, NodePublicKey, prevBlock)
	log.Info("Whether it is the current packaging node", nodePosition, "current node id:", currentCandidateNode.ID, "tcpaddress:", currentCandidateNode.TcpAddress)
	if nodePosition {
		btc := protocols.NewBlockTimeCounter()
		st := time.Now()

		dtx := DelayedTx{
			privateKey: NodePrivateKey,
			publicKey:  NodePublicKey,
			logger:     d.logger,
			time:       st.Unix(),
		}

		_, endTime, err := btc.RangeByTime(st)
		if err != nil {
			log.WithFields(log.Fields{"type": consts.TimeCalcError, "error": err}).Error("on getting end time of generation")
			return err
		}
		done := time.After(endTime.Sub(st))
		txs, err := dtx.RunForDelayBlockID(prevBlock.BlockID + 1)
		if err != nil {
			return err
		}
		//trs, err := processTransactionsNew(d.logger, txs, done, st.Unix())
		trs, err := processTransactions(d.logger, txs, done, st.Unix())
		if err != nil {
			return err
		}
		// Block generation will be started only if we have transactions
		if len(trs) == 0 {
			return nil
		}
		candidateNodesByte, _ := json.Marshal(candidateNodes)
		header := &types.BlockData{
			BlockID:        prevBlock.BlockID + 1,
			Time:           st.Unix(),
			EcosystemID:    0,
			KeyID:          conf.Config.KeyID,
			NodePosition:   currentCandidateNode.ID,
			Version:        consts.BlockVersion,
			ConsensusMode:  consts.CandidateNodeMode,
			CandidateNodes: candidateNodesByte,
		}
		pb := &types.BlockData{
			BlockID:       prevBlock.BlockID,
			Hash:          prevBlock.Hash,
			RollbacksHash: prevBlock.RollbacksHash,
		}

		err = generateCommon(header, pb, trs, NodePrivateKey)
		if err != nil {
			return err
		}
	} else {
		d.sleepTime = 4 * time.Second
		d.logger.WithFields(log.Fields{"type": consts.JustWaiting, "error": err}).Debug("we are not honor node, sleep for 10 seconds")
		return nil
	}
	return nil
}

func generateCommon(blockHeader, prevBlock *types.BlockData, trs []*sqldb.Transaction, key string) error {
	//blockBin, err := generateNextBlockNew(blockHeader, prevBlock, trs)
	blockBin, err := generateNextBlock(blockHeader, prevBlock, trs)
	if err != nil {
		return err
	}
	//err = block.InsertBlockWOForksNew(blockBin, true, false)
	err = block.InsertBlockWOForks(blockBin, true, false)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("on inserting new block")
		return err
	}
	log.WithFields(log.Fields{"block": blockHeader.String(), "type": consts.SyncProcess}).Debug("Generated block ID")
	return nil
}

func GetThisNodePosition(candidateNodes []sqldb.CandidateNode, NodePublicKey string, prevBlock *sqldb.InfoBlock) (sqldb.CandidateNode, bool) {
	candidateNode := sqldb.CandidateNode{}
	if len(candidateNodes) == 0 {

		firstBlock, err := syspar.GetFirstBlockData()
		if err != nil {
			return candidateNode, false
		}
		nodePubKey := syspar.GetNodePubKey()
		if bytes.Equal(firstBlock.NodePublicKey, nodePubKey) {
			candidateNode.ID = 0
			candidateNode.NodePubKey = hex.EncodeToString(nodePubKey)
			syspar.SetRunModel(consts.HonorNodeMode)
			return candidateNode, true
		}
		return candidateNode, false
	}

	if len(candidateNodes) == 1 {
		nodePubKey := candidateNodes[0].NodePubKey
		pk, err := hex.DecodeString(nodePubKey)
		if err != nil {
			return candidateNode, false
		}
		pk = crypto.CutPub(pk)
		if err != nil {
			log.WithFields(log.Fields{"type": consts.ConversionError, "error": err}).Error("decoding node private key from hex")
			return candidateNode, false
		}
		if bytes.Equal(pk, syspar.GetNodePubKey()) {
			return candidateNodes[0], true
		}
		return candidateNode, false
	}

	if len(candidateNodes) == 2 {
		_, NodePublicKey := utils.GetNodeKeys()
		NodePublicKey = "04" + NodePublicKey
		var (
			packageNode sqldb.CandidateNode
			flag        bool
		)
		for _, node := range candidateNodes {

			if NodePublicKey == node.NodePubKey && prevBlock.NodePosition != strconv.FormatInt(node.ID, 10) {
				flag = true
				packageNode = node
				break
			}
			packageNode = node
		}
		if flag {
			return packageNode, flag
		}
		return packageNode, flag
	}
	if len(candidateNodes) > 2 {
		candidateNodesSqrt := math.Sqrt(float64(len(candidateNodes)))
		candidateNodesCeil := math.Ceil(candidateNodesSqrt)
		startBlockId := prevBlock.BlockID - int64(candidateNodesCeil)
		subBlocks, err := sqldb.GetBlockchain(startBlockId, prevBlock.BlockID, sqldb.OrderASC)
		if err != nil {
			log.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("getting recent block")
			return candidateNode, false
		}
		size := len(candidateNodes)
		for i, subBlock := range subBlocks {
			for j := 0; j < size; j++ {
				if candidateNodes[j].ID == subBlock.NodePosition {
					candidateNodes = append(candidateNodes[:j], candidateNodes[j+1:]...)
					size = len(candidateNodes)
					i--
				}
			}
		}
		if len(candidateNodes) > 0 {
			maxVal := candidateNodes[0].ReplyCount
			maxIndex := 0
			for i, node := range candidateNodes {
				if maxVal < node.ReplyCount {
					maxVal = candidateNodes[i].ReplyCount
					maxIndex = i
				}
			}
			_, NodePublicKey := utils.GetNodeKeys()
			if len(NodePublicKey) < 1 {
				log.WithFields(log.Fields{"type": consts.EmptyObject}).Error("node public key is empty")
				return candidateNode, false
			}
			NodePublicKey = "04" + NodePublicKey

			if NodePublicKey == candidateNodes[maxIndex].NodePubKey {
				return candidateNodes[maxIndex], true
			}

		}
	}
	return candidateNode, false
}

func checkVotes(candidateNodes int64, votes int64) bool {
	lessVotes := math.Ceil(float64(candidateNodes / 2))
	if votes >= int64(lessVotes) {
		return true
	}
	return false
}