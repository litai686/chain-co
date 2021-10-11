package proxy

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/go-state-types/dline"
	"github.com/filecoin-project/go-state-types/network"
	api1 "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	miner1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	"github.com/ipfs-force-community/chain-co/api"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

var _ ProxyAPI = (*Proxy)(nil)

type ProxyAPI interface {
	api.Proxy
}

type Proxy struct {
	Select func() (ProxyAPI, error)
}

// impl api.Proxy
func (p *Proxy) BeaconGetEntry(in0 context.Context, in1 abi.ChainEpoch) (out0 *types.BeaconEntry, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api BeaconGetEntry %v", err)
		return
	}
	return cli.BeaconGetEntry(in0, in1)
}

func (p *Proxy) ChainGetBlock(in0 context.Context, in1 cid.Cid) (out0 *types.BlockHeader, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetBlock %v", err)
		return
	}
	return cli.ChainGetBlock(in0, in1)
}

func (p *Proxy) ChainGetBlockMessages(in0 context.Context, in1 cid.Cid) (out0 *api1.BlockMessages, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetBlockMessages %v", err)
		return
	}
	return cli.ChainGetBlockMessages(in0, in1)
}

func (p *Proxy) ChainGetGenesis(in0 context.Context) (out0 *types.TipSet, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetGenesis %v", err)
		return
	}
	return cli.ChainGetGenesis(in0)
}

func (p *Proxy) ChainGetMessage(in0 context.Context, in1 cid.Cid) (out0 *types.Message, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetMessage %v", err)
		return
	}
	return cli.ChainGetMessage(in0, in1)
}

func (p *Proxy) ChainGetMessagesInTipset(in0 context.Context, in1 types.TipSetKey) (out0 []api1.Message, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetMessagesInTipset %v", err)
		return
	}
	return cli.ChainGetMessagesInTipset(in0, in1)
}

func (p *Proxy) ChainGetParentMessages(in0 context.Context, in1 cid.Cid) (out0 []api1.Message, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetParentMessages %v", err)
		return
	}
	return cli.ChainGetParentMessages(in0, in1)
}

func (p *Proxy) ChainGetParentReceipts(in0 context.Context, in1 cid.Cid) (out0 []*types.MessageReceipt, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetParentReceipts %v", err)
		return
	}
	return cli.ChainGetParentReceipts(in0, in1)
}

func (p *Proxy) ChainGetPath(in0 context.Context, in1 types.TipSetKey, in2 types.TipSetKey) (out0 []*api1.HeadChange, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetPath %v", err)
		return
	}
	return cli.ChainGetPath(in0, in1, in2)
}

func (p *Proxy) ChainGetRandomnessFromBeacon(in0 context.Context, in1 types.TipSetKey, in2 crypto.DomainSeparationTag, in3 abi.ChainEpoch, in4 []uint8) (out0 abi.Randomness, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetRandomnessFromBeacon %v", err)
		return
	}
	return cli.ChainGetRandomnessFromBeacon(in0, in1, in2, in3, in4)
}

func (p *Proxy) ChainGetRandomnessFromTickets(in0 context.Context, in1 types.TipSetKey, in2 crypto.DomainSeparationTag, in3 abi.ChainEpoch, in4 []uint8) (out0 abi.Randomness, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetRandomnessFromTickets %v", err)
		return
	}
	return cli.ChainGetRandomnessFromTickets(in0, in1, in2, in3, in4)
}

func (p *Proxy) ChainGetTipSet(in0 context.Context, in1 types.TipSetKey) (out0 *types.TipSet, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetTipSet %v", err)
		return
	}
	return cli.ChainGetTipSet(in0, in1)
}

func (p *Proxy) ChainGetTipSetAfterHeight(in0 context.Context, in1 abi.ChainEpoch, in2 types.TipSetKey) (out0 *types.TipSet, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetTipSetAfterHeight %v", err)
		return
	}
	return cli.ChainGetTipSetAfterHeight(in0, in1, in2)
}

func (p *Proxy) ChainGetTipSetByHeight(in0 context.Context, in1 abi.ChainEpoch, in2 types.TipSetKey) (out0 *types.TipSet, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainGetTipSetByHeight %v", err)
		return
	}
	return cli.ChainGetTipSetByHeight(in0, in1, in2)
}

func (p *Proxy) ChainHasObj(in0 context.Context, in1 cid.Cid) (out0 bool, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainHasObj %v", err)
		return
	}
	return cli.ChainHasObj(in0, in1)
}

func (p *Proxy) ChainHead(in0 context.Context) (out0 *types.TipSet, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainHead %v", err)
		return
	}
	return cli.ChainHead(in0)
}

func (p *Proxy) ChainReadObj(in0 context.Context, in1 cid.Cid) (out0 []uint8, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainReadObj %v", err)
		return
	}
	return cli.ChainReadObj(in0, in1)
}

func (p *Proxy) ChainStatObj(in0 context.Context, in1 cid.Cid, in2 cid.Cid) (out0 api1.ObjStat, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainStatObj %v", err)
		return
	}
	return cli.ChainStatObj(in0, in1, in2)
}

func (p *Proxy) ChainTipSetWeight(in0 context.Context, in1 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api ChainTipSetWeight %v", err)
		return
	}
	return cli.ChainTipSetWeight(in0, in1)
}

func (p *Proxy) GasBatchEstimateMessageGas(in0 context.Context, in1 []*api1.EstimateMessage, in2 uint64, in3 types.TipSetKey) (out0 []*api1.EstimateResult, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api GasBatchEstimateMessageGas %v", err)
		return
	}
	return cli.GasBatchEstimateMessageGas(in0, in1, in2, in3)
}

func (p *Proxy) GasEstimateFeeCap(in0 context.Context, in1 *types.Message, in2 int64, in3 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api GasEstimateFeeCap %v", err)
		return
	}
	return cli.GasEstimateFeeCap(in0, in1, in2, in3)
}

func (p *Proxy) GasEstimateGasLimit(in0 context.Context, in1 *types.Message, in2 types.TipSetKey) (out0 int64, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api GasEstimateGasLimit %v", err)
		return
	}
	return cli.GasEstimateGasLimit(in0, in1, in2)
}

func (p *Proxy) GasEstimateGasPremium(in0 context.Context, in1 uint64, in2 address.Address, in3 int64, in4 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api GasEstimateGasPremium %v", err)
		return
	}
	return cli.GasEstimateGasPremium(in0, in1, in2, in3, in4)
}

func (p *Proxy) GasEstimateMessageGas(in0 context.Context, in1 *types.Message, in2 *api1.MessageSendSpec, in3 types.TipSetKey) (out0 *types.Message, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api GasEstimateMessageGas %v", err)
		return
	}
	return cli.GasEstimateMessageGas(in0, in1, in2, in3)
}

func (p *Proxy) MinerCreateBlock(in0 context.Context, in1 *api1.BlockTemplate) (out0 *types.BlockMsg, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api MinerCreateBlock %v", err)
		return
	}
	return cli.MinerCreateBlock(in0, in1)
}

func (p *Proxy) MinerGetBaseInfo(in0 context.Context, in1 address.Address, in2 abi.ChainEpoch, in3 types.TipSetKey) (out0 *api1.MiningBaseInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api MinerGetBaseInfo %v", err)
		return
	}
	return cli.MinerGetBaseInfo(in0, in1, in2, in3)
}

func (p *Proxy) MpoolPublishByAddr(in0 context.Context, in1 address.Address) (err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api MpoolPublishByAddr %v", err)
		return
	}
	return cli.MpoolPublishByAddr(in0, in1)
}

func (p *Proxy) MpoolPublishMessage(in0 context.Context, in1 *types.SignedMessage) (err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api MpoolPublishMessage %v", err)
		return
	}
	return cli.MpoolPublishMessage(in0, in1)
}

func (p *Proxy) MpoolPush(in0 context.Context, in1 *types.SignedMessage) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api MpoolPush %v", err)
		return
	}
	return cli.MpoolPush(in0, in1)
}

func (p *Proxy) MpoolPushMessage(in0 context.Context, in1 *types.Message, in2 *api1.MessageSendSpec) (out0 *types.SignedMessage, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api MpoolPushMessage %v", err)
		return
	}
	return cli.MpoolPushMessage(in0, in1, in2)
}

func (p *Proxy) MpoolSelect(in0 context.Context, in1 types.TipSetKey, in2 float64) (out0 []*types.SignedMessage, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api MpoolSelect %v", err)
		return
	}
	return cli.MpoolSelect(in0, in1, in2)
}

func (p *Proxy) MpoolSelects(in0 context.Context, in1 types.TipSetKey, in2 []float64) (out0 [][]*types.SignedMessage, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api MpoolSelects %v", err)
		return
	}
	return cli.MpoolSelects(in0, in1, in2)
}

func (p *Proxy) StateAccountKey(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateAccountKey %v", err)
		return
	}
	return cli.StateAccountKey(in0, in1, in2)
}

func (p *Proxy) StateAllMinerFaults(in0 context.Context, in1 abi.ChainEpoch, in2 types.TipSetKey) (out0 []*api1.Fault, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateAllMinerFaults %v", err)
		return
	}
	return cli.StateAllMinerFaults(in0, in1, in2)
}

func (p *Proxy) StateCall(in0 context.Context, in1 *types.Message, in2 types.TipSetKey) (out0 *api1.InvocResult, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateCall %v", err)
		return
	}
	return cli.StateCall(in0, in1, in2)
}

func (p *Proxy) StateChangedActors(in0 context.Context, in1 cid.Cid, in2 cid.Cid) (out0 map[string]types.Actor, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateChangedActors %v", err)
		return
	}
	return cli.StateChangedActors(in0, in1, in2)
}

func (p *Proxy) StateCirculatingSupply(in0 context.Context, in1 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateCirculatingSupply %v", err)
		return
	}
	return cli.StateCirculatingSupply(in0, in1)
}

func (p *Proxy) StateDealProviderCollateralBounds(in0 context.Context, in1 abi.PaddedPieceSize, in2 bool, in3 types.TipSetKey) (out0 api1.DealCollateralBounds, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateDealProviderCollateralBounds %v", err)
		return
	}
	return cli.StateDealProviderCollateralBounds(in0, in1, in2, in3)
}

func (p *Proxy) StateGetActor(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *types.Actor, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateGetActor %v", err)
		return
	}
	return cli.StateGetActor(in0, in1, in2)
}

func (p *Proxy) StateGetRandomnessFromBeacon(in0 context.Context, in1 crypto.DomainSeparationTag, in2 abi.ChainEpoch, in3 []uint8, in4 types.TipSetKey) (out0 abi.Randomness, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateGetRandomnessFromBeacon %v", err)
		return
	}
	return cli.StateGetRandomnessFromBeacon(in0, in1, in2, in3, in4)
}

func (p *Proxy) StateGetRandomnessFromTickets(in0 context.Context, in1 crypto.DomainSeparationTag, in2 abi.ChainEpoch, in3 []uint8, in4 types.TipSetKey) (out0 abi.Randomness, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateGetRandomnessFromTickets %v", err)
		return
	}
	return cli.StateGetRandomnessFromTickets(in0, in1, in2, in3, in4)
}

func (p *Proxy) StateListActors(in0 context.Context, in1 types.TipSetKey) (out0 []address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateListActors %v", err)
		return
	}
	return cli.StateListActors(in0, in1)
}

func (p *Proxy) StateListMiners(in0 context.Context, in1 types.TipSetKey) (out0 []address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateListMiners %v", err)
		return
	}
	return cli.StateListMiners(in0, in1)
}

func (p *Proxy) StateLookupID(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateLookupID %v", err)
		return
	}
	return cli.StateLookupID(in0, in1, in2)
}

func (p *Proxy) StateMarketBalance(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 api1.MarketBalance, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMarketBalance %v", err)
		return
	}
	return cli.StateMarketBalance(in0, in1, in2)
}

func (p *Proxy) StateMarketDeals(in0 context.Context, in1 types.TipSetKey) (out0 map[string]api1.MarketDeal, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMarketDeals %v", err)
		return
	}
	return cli.StateMarketDeals(in0, in1)
}

func (p *Proxy) StateMarketParticipants(in0 context.Context, in1 types.TipSetKey) (out0 map[string]api1.MarketBalance, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMarketParticipants %v", err)
		return
	}
	return cli.StateMarketParticipants(in0, in1)
}

func (p *Proxy) StateMarketStorageDeal(in0 context.Context, in1 abi.DealID, in2 types.TipSetKey) (out0 *api1.MarketDeal, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMarketStorageDeal %v", err)
		return
	}
	return cli.StateMarketStorageDeal(in0, in1, in2)
}

func (p *Proxy) StateMinerActiveSectors(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 []*miner.SectorOnChainInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerActiveSectors %v", err)
		return
	}
	return cli.StateMinerActiveSectors(in0, in1, in2)
}

func (p *Proxy) StateMinerAvailableBalance(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerAvailableBalance %v", err)
		return
	}
	return cli.StateMinerAvailableBalance(in0, in1, in2)
}

func (p *Proxy) StateMinerDeadlines(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 []api1.Deadline, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerDeadlines %v", err)
		return
	}
	return cli.StateMinerDeadlines(in0, in1, in2)
}

func (p *Proxy) StateMinerFaults(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 bitfield.BitField, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerFaults %v", err)
		return
	}
	return cli.StateMinerFaults(in0, in1, in2)
}

func (p *Proxy) StateMinerInfo(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 miner.MinerInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerInfo %v", err)
		return
	}
	return cli.StateMinerInfo(in0, in1, in2)
}

func (p *Proxy) StateMinerInitialPledgeCollateral(in0 context.Context, in1 address.Address, in2 miner1.SectorPreCommitInfo, in3 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerInitialPledgeCollateral %v", err)
		return
	}
	return cli.StateMinerInitialPledgeCollateral(in0, in1, in2, in3)
}

func (p *Proxy) StateMinerPartitions(in0 context.Context, in1 address.Address, in2 uint64, in3 types.TipSetKey) (out0 []api1.Partition, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerPartitions %v", err)
		return
	}
	return cli.StateMinerPartitions(in0, in1, in2, in3)
}

func (p *Proxy) StateMinerPower(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *api1.MinerPower, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerPower %v", err)
		return
	}
	return cli.StateMinerPower(in0, in1, in2)
}

func (p *Proxy) StateMinerPreCommitDepositForPower(in0 context.Context, in1 address.Address, in2 miner1.SectorPreCommitInfo, in3 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerPreCommitDepositForPower %v", err)
		return
	}
	return cli.StateMinerPreCommitDepositForPower(in0, in1, in2, in3)
}

func (p *Proxy) StateMinerProvingDeadline(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *dline.Info, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerProvingDeadline %v", err)
		return
	}
	return cli.StateMinerProvingDeadline(in0, in1, in2)
}

func (p *Proxy) StateMinerRecoveries(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 bitfield.BitField, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerRecoveries %v", err)
		return
	}
	return cli.StateMinerRecoveries(in0, in1, in2)
}

func (p *Proxy) StateMinerSectorAllocated(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 bool, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerSectorAllocated %v", err)
		return
	}
	return cli.StateMinerSectorAllocated(in0, in1, in2, in3)
}

func (p *Proxy) StateMinerSectorCount(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 api1.MinerSectors, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerSectorCount %v", err)
		return
	}
	return cli.StateMinerSectorCount(in0, in1, in2)
}

func (p *Proxy) StateMinerSectors(in0 context.Context, in1 address.Address, in2 *bitfield.BitField, in3 types.TipSetKey) (out0 []*miner.SectorOnChainInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateMinerSectors %v", err)
		return
	}
	return cli.StateMinerSectors(in0, in1, in2, in3)
}

func (p *Proxy) StateNetworkName(in0 context.Context) (out0 dtypes.NetworkName, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateNetworkName %v", err)
		return
	}
	return cli.StateNetworkName(in0)
}

func (p *Proxy) StateNetworkVersion(in0 context.Context, in1 types.TipSetKey) (out0 network.Version, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateNetworkVersion %v", err)
		return
	}
	return cli.StateNetworkVersion(in0, in1)
}

func (p *Proxy) StateReadState(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *api1.ActorState, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateReadState %v", err)
		return
	}
	return cli.StateReadState(in0, in1, in2)
}

func (p *Proxy) StateSearchMsg(in0 context.Context, in1 types.TipSetKey, in2 cid.Cid, in3 abi.ChainEpoch, in4 bool) (out0 *api1.MsgLookup, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateSearchMsg %v", err)
		return
	}
	return cli.StateSearchMsg(in0, in1, in2, in3, in4)
}

func (p *Proxy) StateSectorExpiration(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 *miner.SectorExpiration, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateSectorExpiration %v", err)
		return
	}
	return cli.StateSectorExpiration(in0, in1, in2, in3)
}

func (p *Proxy) StateSectorGetInfo(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 *miner.SectorOnChainInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateSectorGetInfo %v", err)
		return
	}
	return cli.StateSectorGetInfo(in0, in1, in2, in3)
}

func (p *Proxy) StateSectorPartition(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 *miner.SectorLocation, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateSectorPartition %v", err)
		return
	}
	return cli.StateSectorPartition(in0, in1, in2, in3)
}

func (p *Proxy) StateSectorPreCommitInfo(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 miner.SectorPreCommitOnChainInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateSectorPreCommitInfo %v", err)
		return
	}
	return cli.StateSectorPreCommitInfo(in0, in1, in2, in3)
}

func (p *Proxy) StateVMCirculatingSupplyInternal(in0 context.Context, in1 types.TipSetKey) (out0 api1.CirculatingSupply, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateVMCirculatingSupplyInternal %v", err)
		return
	}
	return cli.StateVMCirculatingSupplyInternal(in0, in1)
}

func (p *Proxy) StateVerifiedClientStatus(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateVerifiedClientStatus %v", err)
		return
	}
	return cli.StateVerifiedClientStatus(in0, in1, in2)
}

func (p *Proxy) StateVerifiedRegistryRootKey(in0 context.Context, in1 types.TipSetKey) (out0 address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateVerifiedRegistryRootKey %v", err)
		return
	}
	return cli.StateVerifiedRegistryRootKey(in0, in1)
}

func (p *Proxy) StateVerifierStatus(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateVerifierStatus %v", err)
		return
	}
	return cli.StateVerifierStatus(in0, in1, in2)
}

func (p *Proxy) StateWaitMsg(in0 context.Context, in1 cid.Cid, in2 uint64, in3 abi.ChainEpoch, in4 bool) (out0 *api1.MsgLookup, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api StateWaitMsg %v", err)
		return
	}
	return cli.StateWaitMsg(in0, in1, in2, in3, in4)
}

func (p *Proxy) SyncState(in0 context.Context) (out0 *api1.SyncState, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api SyncState %v", err)
		return
	}
	return cli.SyncState(in0)
}

func (p *Proxy) SyncSubmitBlock(in0 context.Context, in1 *types.BlockMsg) (err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api SyncSubmitBlock %v", err)
		return
	}
	return cli.SyncSubmitBlock(in0, in1)
}

func (p *Proxy) Version(in0 context.Context) (out0 api1.APIVersion, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api Version %v", err)
		return
	}
	return cli.Version(in0)
}

func (p *Proxy) WalletBalance(in0 context.Context, in1 address.Address) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api WalletBalance %v", err)
		return
	}
	return cli.WalletBalance(in0, in1)
}

func (p *Proxy) WalletHas(in0 context.Context, in1 address.Address) (out0 bool, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api WalletHas %v", err)
		return
	}
	return cli.WalletHas(in0, in1)
}

func (p *Proxy) WalletSign(in0 context.Context, in1 address.Address, in2 []uint8) (out0 *crypto.Signature, err error) {
	cli, err := p.Select()
	if err != nil {
		err = xerrors.Errorf("api WalletSign %v", err)
		return
	}
	return cli.WalletSign(in0, in1, in2)
}
