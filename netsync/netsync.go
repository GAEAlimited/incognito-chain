package netsync

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/incognitochain/incognito-chain/blockchain"
	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/mempool"
	"github.com/incognitochain/incognito-chain/metadata"
	"github.com/incognitochain/incognito-chain/peer"
	"github.com/incognitochain/incognito-chain/pubsub"
	"github.com/incognitochain/incognito-chain/wire"
	libp2p "github.com/libp2p/go-libp2p-peer"
	"github.com/patrickmn/go-cache"
)

// NetSync is a gate for message to enter node from network (after Peerconn),
// all message must be process by NetSync before proccessed by other package in node
// NetSync parses message from other peer, identifies type of message
// After parsing, it will detect if message is duplicate or not
// If message is duplicate then discard it, otherwise pass it to the right handler
// NetSync start when node start and run all the time while node is alive
// and it will stop when received quit signal
type NetSync struct {
	started  int32
	shutdown int32

	cMessage chan interface{}
	cQuit    chan struct{}

	config *NetSyncConfig
	cache  *NetSyncCache
}

type NetSyncConfig struct {
	BlockChain            *blockchain.BlockChain
	ChainParam            *blockchain.Params
	TxMemPool             *mempool.TxPool
	ShardToBeaconPool     blockchain.ShardToBeaconPool
	CrossShardPool        map[byte]blockchain.CrossShardPool
	PubSubManager         *pubsub.PubSubManager
	TransactionEvent      pubsub.EventChannel // transaction event
	RoleInCommitteesEvent pubsub.EventChannel // role in committees event
	BeaconBlockEvent      pubsub.EventChannel // beacon block event
	ShardBlockEvent       pubsub.EventChannel // shard block event
	RelayShard            []byte
	RoleInCommittees      int
	roleInCommitteesMtx   sync.RWMutex
	Server                interface {
		// list functions callback which are assigned from Server struct
		PushMessageToPeer(wire.Message, libp2p.ID) error
		PushMessageToAll(wire.Message) error
	}
	Consensus interface {
		OnBFTMsg(wire.Message)
	}
}

type NetSyncCache struct {
	blockCache    *cache.Cache
	txCache       *cache.Cache
	txCacheMtx    sync.Mutex
	blockCacheMtx sync.Mutex
}

func (netSync *NetSync) Init(cfg *NetSyncConfig) {
	netSync.config = cfg
	netSync.cQuit = make(chan struct{})
	netSync.cMessage = make(chan interface{}, 1000)

	// init cache
	blockCache := cache.New(messageLiveTime, messageCleanupInterval)
	txCache := cache.New(messageLiveTime, messageCleanupInterval)
	netSync.cache = &NetSyncCache{
		txCache:    txCache,
		blockCache: blockCache,
	}

	// register pubsub channel
	_, subChanTx, err := netSync.config.PubSubManager.RegisterNewSubscriber(pubsub.TransactionHashEnterNodeTopic)
	if err != nil {
		Logger.log.Error(err)
	}
	netSync.config.TransactionEvent = subChanTx
	_, subChanRole, err := netSync.config.PubSubManager.RegisterNewSubscriber(pubsub.ShardRoleTopic)
	if err != nil {
		Logger.log.Error(err)
	}
	netSync.config.RoleInCommitteesEvent = subChanRole
	_, subChanBeaconBlock, err := netSync.config.PubSubManager.RegisterNewSubscriber(pubsub.NewBeaconBlockTopic)
	if err != nil {
		Logger.log.Error(err)
	}
	netSync.config.BeaconBlockEvent = subChanBeaconBlock
	_, subChanShardBlock, err := netSync.config.PubSubManager.RegisterNewSubscriber(pubsub.NewShardblockTopic)
	if err != nil {
		Logger.log.Error(err)
	}
	netSync.config.ShardBlockEvent = subChanShardBlock
}

func (netSync *NetSync) Start() error {
	// Already started?
	if atomic.AddInt32(&netSync.started, 1) != 1 {
		return NewNetSyncError(AlreadyStartError, errors.New("Already started"))
	}
	Logger.log.Debug("Starting sync manager")
	//netSync.waitgroup.Add(1)
	go netSync.messageHandler()
	go netSync.cacheLoop()
	return nil
}

// Stop gracefully shuts down the sync manager by stopping all asynchronous
// handlers and waiting for them to finish.
func (netSync *NetSync) Stop() {
	if atomic.AddInt32(&netSync.shutdown, 1) != 1 {
		Logger.log.Warn("Sync manager is already in the process of shutting down")
	}

	Logger.log.Warn("Sync manager shutting down")
	close(netSync.cQuit)
}

// messageHandler is the main handler for the sync manager.  It must be run as a
// goroutine.  It processes block and inv messages in a separate goroutine
// from the peer handlers so the block (MsgBlock) messages are handled by a
// single thread without needing to lock memory data structures.  This is
// important because the sync manager controls which blocks are needed and how
// the fetching should proceed.
func (netSync *NetSync) messageHandler() {
out:
	for {
		select {
		case msgChan := <-netSync.cMessage:
			{
				go func(msgC interface{}) {
					switch msg := msgC.(type) {
					case *wire.MessageTx, *wire.MessageTxToken, *wire.MessageTxPrivacyToken:
						{
							switch msg := msgC.(type) {
							case *wire.MessageTx:
								{
									netSync.HandleMessageTx(msg)
								}
							case *wire.MessageTxToken:
								{
									netSync.HandleMessageTxToken(msg)
								}
							case *wire.MessageTxPrivacyToken:
								{
									netSync.HandleMessageTxPrivacyToken(msg)
								}
							}
						}
					case *wire.MessageBFTPropose:
						{
							netSync.HandleMessageBFTMsg(msg)
						}
					case *wire.MessageBFTAgree:
						{
							netSync.HandleMessageBFTMsg(msg)
						}
					case *wire.MessageBFTCommit:
						{
							netSync.HandleMessageBFTMsg(msg)
						}
					case *wire.MessageBFTReady:
						{
							netSync.HandleMessageBFTMsg(msg)
						}
					case *wire.MessageBFTReq:
						{
							netSync.HandleMessageBFTMsg(msg)
						}
					case *wire.MessageBlockBeacon:
						{
							netSync.HandleMessageBeaconBlock(msg)
						}
					case *wire.MessageBlockShard:
						{
							netSync.HandleMessageShardBlock(msg)
						}
					case *wire.MessageGetCrossShard:
						{
							netSync.HandleMessageGetCrossShard(msg)
						}
					case *wire.MessageCrossShard:
						{
							netSync.HandleMessageCrossShard(msg)
						}
					case *wire.MessageGetShardToBeacon:
						{
							netSync.HandleMessageGetShardToBeacon(msg)
						}
					case *wire.MessageShardToBeacon:
						{
							netSync.HandleMessageShardToBeacon(msg)
						}
					case *wire.MessageGetBlockBeacon:
						{
							netSync.HandleMessageGetBlockBeacon(msg)
						}
					case *wire.MessageGetBlockShard:
						{
							netSync.HandleMessageGetBlockShard(msg)
						}
					case *wire.MessagePeerState:
						{
							netSync.HandleMessagePeerState(msg)
						}
					default:
						Logger.log.Debugf("Invalid message type in block "+"handler: %T", msg)
					}
				}(msgChan)
			}
		case msgChan := <-netSync.cQuit:
			{
				Logger.log.Warn(msgChan)
				break out
			}
		}
	}

	//netSync.waitgroup.Done()
	close(netSync.cMessage)
	Logger.log.Debug("Block handler done")
}

func (netSync *NetSync) QueueTx(peer *peer.Peer, msg *wire.MessageTx, done chan struct{}) error {
	// Don't accept more transactions if we're shutting down.
	if atomic.LoadInt32(&netSync.shutdown) != 0 {
		done <- struct{}{}
		return NewNetSyncError(AlreadyShutdownError, errors.New("We're shutting down"))
	}
	netSync.cMessage <- msg
	return nil
}

func (netSync *NetSync) QueueTxToken(peer *peer.Peer, msg *wire.MessageTxToken, done chan struct{}) error {
	// Don't accept more transactions if we're shutting down.
	if atomic.LoadInt32(&netSync.shutdown) != 0 {
		done <- struct{}{}
		return NewNetSyncError(AlreadyShutdownError, errors.New("we're shutting down"))
	}
	netSync.cMessage <- msg
	return nil
}

func (netSync *NetSync) QueueTxPrivacyToken(peer *peer.Peer, msg *wire.MessageTxPrivacyToken, done chan struct{}) error {
	// Don't accept more transactions if we're shutting down.
	if atomic.LoadInt32(&netSync.shutdown) != 0 {
		done <- struct{}{}
		return NewNetSyncError(AlreadyShutdownError, errors.New("We're shutting down"))
	}
	netSync.cMessage <- msg
	return nil
}

// handleTxMsg handles transaction messages from all peers.
func (netSync *NetSync) HandleMessageTx(msg *wire.MessageTx) {
	Logger.log.Debug("Handling new message tx")
	if !netSync.HandleTxWithRole(msg.Transaction) {
		return
	}
	if isAdded := netSync.HandleCacheTx(*msg.Transaction.Hash()); !isAdded {
		hash, _, err := netSync.config.TxMemPool.MaybeAcceptTransaction(msg.Transaction)
		if err != nil {
			Logger.log.Error(err)
		} else {
			// Broadcast to network
			Logger.log.Debugf("there is hash of transaction %s", hash.String())
			err := netSync.config.Server.PushMessageToAll(msg)
			if err != nil {
				Logger.log.Error(err)
			} else {
				netSync.config.TxMemPool.MarkForwardedTransaction(*msg.Transaction.Hash())
			}
		}
	}
	Logger.log.Debug("Transaction %+v found in cache", *msg.Transaction.Hash())
}

// handleTxMsg handles transaction messages from all peers.
func (netSync *NetSync) HandleMessageTxToken(msg *wire.MessageTxToken) {
	Logger.log.Debug("Handling new message tx")
	if !netSync.HandleTxWithRole(msg.Transaction) {
		return
	}
	if isAdded := netSync.HandleCacheTx(*msg.Transaction.Hash()); !isAdded {
		hash, _, err := netSync.config.TxMemPool.MaybeAcceptTransaction(msg.Transaction)

		if err != nil {
			Logger.log.Error(err)
		} else {
			Logger.log.Debugf("there is hash of transaction %s", hash.String())
			// Broadcast to network
			err := netSync.config.Server.PushMessageToAll(msg)
			if err != nil {
				Logger.log.Error(err)
			} else {
				netSync.config.TxMemPool.MarkForwardedTransaction(*msg.Transaction.Hash())
			}
		}
	}
	Logger.log.Debug("Transaction %+v found in cache", *msg.Transaction.Hash())
}

// handleTxMsg handles transaction messages from all peers.
func (netSync *NetSync) HandleMessageTxPrivacyToken(msg *wire.MessageTxPrivacyToken) {
	Logger.log.Debug("Handling new message tx")
	if !netSync.HandleTxWithRole(msg.Transaction) {
		return
	}
	if isAdded := netSync.HandleCacheTx(*msg.Transaction.Hash()); !isAdded {
		hash, _, err := netSync.config.TxMemPool.MaybeAcceptTransaction(msg.Transaction)
		if err != nil {
			Logger.log.Error(err)
		} else {
			Logger.log.Debugf("Node got hash of transaction %s", hash.String())
			// Broadcast to network
			err := netSync.config.Server.PushMessageToAll(msg)
			if err != nil {
				Logger.log.Error(err)
			} else {
				netSync.config.TxMemPool.MarkForwardedTransaction(*msg.Transaction.Hash())
			}
		}
	}
	Logger.log.Debug("Transaction %+v found in cache", *msg.Transaction.Hash())
}

// QueueBlock adds the passed block message and peer to the block handling
// queue. Responds to the done channel argument after the block message is
// processed.
func (netSync *NetSync) QueueBlock(_ *peer.Peer, msg wire.Message, done chan struct{}) {
	// Don't accept more transactions if we're shutting down.
	if atomic.LoadInt32(&netSync.shutdown) != 0 {
		done <- struct{}{}
		return
	}
	netSync.cMessage <- msg
}

func (netSync *NetSync) QueueGetBlockShard(peer *peer.Peer, msg *wire.MessageGetBlockShard, done chan struct{}) {
	// Don't accept more transactions if we're shutting down.
	if atomic.LoadInt32(&netSync.shutdown) != 0 {
		done <- struct{}{}
		return
	}
	netSync.cMessage <- msg
}

func (netSync *NetSync) QueueGetBlockBeacon(peer *peer.Peer, msg *wire.MessageGetBlockBeacon, done chan struct{}) {
	// Don't accept more transactions if we're shutting down.
	if atomic.LoadInt32(&netSync.shutdown) != 0 {
		done <- struct{}{}
		return
	}
	netSync.cMessage <- msg
}

func (netSync *NetSync) QueueMessage(peer *peer.Peer, msg wire.Message, done chan struct{}) {
	// Don't accept more transactions if we're shutting down.
	if atomic.LoadInt32(&netSync.shutdown) != 0 {
		done <- struct{}{}
		return
	}
	netSync.cMessage <- msg
}

func (netSync *NetSync) HandleMessageBeaconBlock(msg *wire.MessageBlockBeacon) {
	Logger.log.Debug("Handling new message BlockBeacon")
	//if oldBlock := netSync.IsOldBeaconBlock(msg.Block.Header.Height); !oldBlock {
	if isAdded := netSync.HandleCacheBlock("b" + msg.Block.Header.Hash().String()); !isAdded {
		netSync.config.BlockChain.OnBlockBeaconReceived(msg.Block)
	}
	//}
}

func (netSync *NetSync) HandleMessageShardBlock(msg *wire.MessageBlockShard) {
	Logger.log.Debug("Handling new message BlockShard")
	if isAdded := netSync.HandleCacheBlock("s" + msg.Block.Header.Hash().String()); !isAdded {
		netSync.config.BlockChain.OnBlockShardReceived(msg.Block)
		return
	}
}

func (netSync *NetSync) HandleMessageCrossShard(msg *wire.MessageCrossShard) {
	Logger.log.Debug("Handling new message CrossShard")
	if isAdded := netSync.HandleCacheBlock("c" + msg.Block.Header.Hash().String()); !isAdded {
		netSync.config.BlockChain.OnCrossShardBlockReceived(msg.Block)
	}

}

func (netSync *NetSync) HandleMessageShardToBeacon(msg *wire.MessageShardToBeacon) {
	Logger.log.Debug("Handling new message ShardToBeacon")
	if isAdded := netSync.HandleCacheBlock("s2b" + msg.Block.Header.Hash().String()); !isAdded {
		netSync.config.BlockChain.OnShardToBeaconBlockReceived(msg.Block)
	}
}

func (netSync *NetSync) HandleMessageBFTMsg(msg wire.Message) {
	Logger.log.Debug("Handling new message BFTMsg")
	if err := msg.VerifyMsgSanity(); err != nil {
		Logger.log.Error(err)
		return
	}
	netSync.config.Consensus.OnBFTMsg(msg)
}

func (netSync *NetSync) HandleMessagePeerState(msg *wire.MessagePeerState) {
	Logger.log.Debug("Handling new message peerstate", msg.SenderID)
	peerID, err := libp2p.IDB58Decode(msg.SenderID)
	if err != nil {
		Logger.log.Error(err)
		return
	}
	netSync.config.BlockChain.OnPeerStateReceived(&msg.Beacon, &msg.Shards, &msg.ShardToBeaconPool, &msg.CrossShardPool, peerID)
}

func (netSync *NetSync) HandleMessageGetBlockShard(msg *wire.MessageGetBlockShard) {
	Logger.log.Debug("Handling new message - " + wire.CmdGetBlockShard)
	peerID, err := libp2p.IDB58Decode(msg.SenderID)
	if err != nil {
		Logger.log.Error(err)
		return
	}
	if msg.ByHash {
		netSync.getBlkShardByHashAndSend(peerID, 0, msg.BlkHashes, 0)
	} else {
		netSync.getBlockShardByHeightAndSend(peerID, msg.FromPool, 0, msg.BySpecificHeight, msg.ShardID, msg.BlkHeights, 0)
	}
}

func (netSync *NetSync) HandleMessageGetBlockBeacon(msg *wire.MessageGetBlockBeacon) {
	Logger.log.Debug("Handling new message - " + wire.CmdGetBlockBeacon)
	peerID, err := libp2p.IDB58Decode(msg.SenderID)
	if err != nil {
		Logger.log.Error(err)
		return
	}
	if msg.ByHash {
		netSync.getBlockBeaconByHashAndSend(peerID, msg.BlkHashes)
	} else {
		netSync.getBlockBeaconByHeightAndSend(peerID, msg.FromPool, msg.BySpecificHeight, msg.BlkHeights)
	}
}

func (netSync *NetSync) HandleMessageGetShardToBeacon(msg *wire.MessageGetShardToBeacon) {
	Logger.log.Debug("Handling new message getshardtobeacon")
	peerID, err := libp2p.IDB58Decode(msg.SenderID)
	if err != nil {
		Logger.log.Error(err)
		return
	}
	if msg.ByHash {
		netSync.getBlkShardByHashAndSend(peerID, 2, msg.BlkHashes, 0)
	} else {
		netSync.getBlockShardByHeightAndSend(peerID, msg.FromPool, 2, msg.BySpecificHeight, msg.ShardID, msg.BlkHeights, 0)
	}
}

func (netSync *NetSync) HandleMessageGetCrossShard(msg *wire.MessageGetCrossShard) {
	Logger.log.Debug("Handling new message getcrossshard")
	peerID, err := libp2p.IDB58Decode(msg.SenderID)
	if err != nil {
		Logger.log.Error(err)
		return
	}
	if msg.ByHash {
		netSync.getBlkShardByHashAndSend(peerID, 1, msg.BlkHashes, msg.ToShardID)
	} else {
		netSync.getBlockShardByHeightAndSend(peerID, msg.FromPool, 1, msg.BySpecificHeight, msg.FromShardID, msg.BlkHeights, msg.ToShardID)
	}
}

func (netSync *NetSync) HandleCacheBlock(blockHash string) bool {
	netSync.cache.blockCacheMtx.Lock()
	defer netSync.cache.blockCacheMtx.Unlock()
	_, ok := netSync.cache.blockCache.Get(blockHash)
	if ok {
		return true
	}
	err := netSync.cache.blockCache.Add(blockHash, 1, messageLiveTime)
	if err != nil {
		Logger.log.Error(err)
	}
	return false
}

func (netSync *NetSync) HandleCacheTx(txHash common.Hash) bool {
	netSync.cache.txCacheMtx.Lock()
	defer netSync.cache.txCacheMtx.Unlock()
	_, ok := netSync.cache.txCache.Get(txHash.String())
	if ok {
		return true
	}
	err := netSync.cache.txCache.Add(txHash.String(), 1, messageLiveTime)
	if err != nil {
		Logger.log.Error(err)
	}
	return false
}

func (netSync *NetSync) HandleCacheTxHash(txHash common.Hash) {
	netSync.cache.txCacheMtx.Lock()
	defer netSync.cache.txCacheMtx.Unlock()
	err := netSync.cache.txCache.Add(txHash.String(), 1, messageLiveTime)
	if err != nil {
		Logger.log.Error(err)
	}
}

func (netSync *NetSync) HandleTxWithRole(tx metadata.Transaction) bool {
	senderShardID := common.GetShardIDFromLastByte(tx.GetSenderAddrLastByte())
	for _, shardID := range netSync.config.RelayShard {
		if senderShardID == shardID {
			return true
		}
	}
	netSync.config.roleInCommitteesMtx.RLock()
	if netSync.config.RoleInCommittees > -1 && byte(netSync.config.RoleInCommittees) == senderShardID {
		netSync.config.roleInCommitteesMtx.RUnlock()
		return true
	} else {
		netSync.config.roleInCommitteesMtx.RUnlock()
		return false
	}
}

func (netSync *NetSync) cacheLoop() {
	for w := 0; w < workers; w++ {
		go netSync.HandleCacheTxHashWorker(netSync.config.TransactionEvent)
	}
	for {
		select {
		case msg := <-netSync.config.ShardBlockEvent:
			{
				if shardBlock, ok := msg.Value.(*blockchain.ShardBlock); !ok {
					continue
				} else {
					go netSync.HandleCacheBlock("s" + shardBlock.Header.Hash().String())
				}
			}
		case msg := <-netSync.config.BeaconBlockEvent:
			{
				if beaconBlock, ok := msg.Value.(*blockchain.BeaconBlock); !ok {
					continue
				} else {
					go netSync.HandleCacheBlock("b" + beaconBlock.Header.Hash().String())
				}
			}
		case msg := <-netSync.config.RoleInCommitteesEvent:
			{
				if shardID, ok := msg.Value.(int); !ok {
					continue
				} else {
					netSync.config.roleInCommitteesMtx.Lock()
					netSync.config.RoleInCommittees = shardID
					netSync.config.roleInCommitteesMtx.Unlock()
				}
			}
		}
	}
}

func (netSync *NetSync) HandleCacheTxHashWorker(event pubsub.EventChannel) {
	for msg := range event {
		value, ok := msg.Value.(common.Hash)
		if !ok {
			continue
		}
		go netSync.HandleCacheTxHash(value)
		time.Sleep(time.Nanosecond)
	}
}
