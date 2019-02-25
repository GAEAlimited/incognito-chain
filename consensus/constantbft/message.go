package constantbft

import (
	"encoding/json"

	"github.com/ninjadotorg/constant/cashec"

	"github.com/ninjadotorg/constant/common"

	"github.com/ninjadotorg/constant/blockchain"

	"github.com/ninjadotorg/constant/wire"
)

func (self *Engine) OnBFTMsg(msg wire.Message) {
	self.cBFTMsg <- msg
	return
}

func (self *Engine) OnInvalidBlockReceived(blockHash string, shardID byte, reason string) {
	// leave empty for now
	Logger.log.Error(blockHash, shardID, reason)
	return
}

func MakeMsgBFTReady(bestStateHash common.Hash, round int, userKeySet *cashec.KeySet) (wire.Message, error) {
	msg, err := wire.MakeEmptyMessage(wire.CmdBFTReady)
	if err != nil {
		Logger.log.Error(err)
		return msg, err
	}
	msg.(*wire.MessageBFTReady).BestStateHash = bestStateHash
	msg.(*wire.MessageBFTReady).Round = round
	msg.(*wire.MessageBFTReady).Pubkey = userKeySet.GetPublicKeyB58()
	err = msg.(*wire.MessageBFTReady).SignMsg(userKeySet)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func MakeMsgBFTPropose(block json.RawMessage, layer string, shardID byte, userKeySet *cashec.KeySet) (wire.Message, error) {
	msg, err := wire.MakeEmptyMessage(wire.CmdBFTPropose)
	if err != nil {
		Logger.log.Error(err)
		return msg, err
	}
	msg.(*wire.MessageBFTPropose).Block = block
	msg.(*wire.MessageBFTPropose).Layer = layer
	msg.(*wire.MessageBFTPropose).ShardID = shardID
	msg.(*wire.MessageBFTPropose).Pubkey = userKeySet.GetPublicKeyB58()
	err = msg.(*wire.MessageBFTPropose).SignMsg(userKeySet)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func MakeMsgBFTPrepare(Ri []byte, userKeySet *cashec.KeySet, blkHash common.Hash) (wire.Message, error) {
	msg, err := wire.MakeEmptyMessage(wire.CmdBFTPrepare)
	if err != nil {
		Logger.log.Error(err)

		return msg, err
	}
	msg.(*wire.MessageBFTPrepare).Ri = Ri
	msg.(*wire.MessageBFTPrepare).Pubkey = userKeySet.GetPublicKeyB58()
	msg.(*wire.MessageBFTPrepare).BlkHash = blkHash
	err = msg.(*wire.MessageBFTPrepare).SignMsg(userKeySet)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func MakeMsgBFTCommit(commitSig string, R string, validatorsIdx []int, userKeySet *cashec.KeySet) (wire.Message, error) {
	msg, err := wire.MakeEmptyMessage(wire.CmdBFTCommit)
	if err != nil {
		Logger.log.Error(err)
		return msg, err
	}
	msg.(*wire.MessageBFTCommit).CommitSig = commitSig
	msg.(*wire.MessageBFTCommit).R = R
	msg.(*wire.MessageBFTCommit).ValidatorsIdx = validatorsIdx
	msg.(*wire.MessageBFTCommit).Pubkey = userKeySet.GetPublicKeyB58()
	err = msg.(*wire.MessageBFTCommit).SignMsg(userKeySet)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func MakeMsgBeaconBlock(block *blockchain.BeaconBlock) (wire.Message, error) {
	msg, err := wire.MakeEmptyMessage(wire.CmdBlockBeacon)
	if err != nil {
		Logger.log.Error(err)
		return msg, err
	}
	msg.(*wire.MessageBlockBeacon).Block = *block
	return msg, nil
}

func MakeMsgShardBlock(block *blockchain.ShardBlock) (wire.Message, error) {
	msg, err := wire.MakeEmptyMessage(wire.CmdBlockShard)
	if err != nil {
		Logger.log.Error(err)
		return msg, err
	}
	msg.(*wire.MessageBlockShard).Block = *block
	return msg, nil
}

func MakeMsgShardToBeaconBlock(block *blockchain.ShardToBeaconBlock) (wire.Message, error) {
	msg, err := wire.MakeEmptyMessage(wire.CmdBlkShardToBeacon)
	if err != nil {
		Logger.log.Error(err)
		return msg, err
	}
	msg.(*wire.MessageShardToBeacon).Block = *block
	return msg, nil
}

func MakeMsgCrossShardBlock(block *blockchain.CrossShardBlock) (wire.Message, error) {
	msg, err := wire.MakeEmptyMessage(wire.CmdCrossShard)
	if err != nil {
		Logger.log.Error(err)
		return msg, err
	}
	msg.(*wire.MessageCrossShard).Block = *block
	return msg, nil
}
