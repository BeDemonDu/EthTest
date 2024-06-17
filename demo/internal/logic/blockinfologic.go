package logic

import (
	"EthTest/demo/internal/svc"
	"EthTest/demo/internal/types"
	"context"
	"fmt"
	mtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBlockInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockInfoLogic {
	return &BlockInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type BlockInfo struct {
	Number     int64
	ParentHash string

	Hash string
	Ts   uint64
	Len  int
}

func (l *BlockInfoLogic) BlockInfo(req *types.BlockInfoRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("test1")
	var client *ethclient.Client
	client, err = ethclient.Dial(l.svcCtx.Config.EthClient.BaseUrl) // 8545 是默认的以太坊节点 RPC 端口
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	// 要查询的区块号
	fmt.Println("test2", req.BlockId)
	blockNumber := new(big.Int)
	blockNumber.SetInt64(req.BlockId)

	// 获取指定区块的信息

	var blockInfo *mtypes.Block
	blockInfo, err = client.BlockByNumber(l.ctx, blockNumber)
	if err != nil {
		fmt.Println(err.Error(), l.ctx)
		log.Fatal(err)
		return
	}
	if blockInfo == nil {
		resp.Data = "block empty"
		return
	}
	info := BlockInfo{
		Number:     blockInfo.Number().Int64(),
		ParentHash: blockInfo.ParentHash().Hex(),
		Hash:       blockInfo.Hash().Hex(),
		Ts:         blockInfo.Time(),
		Len:        len(blockInfo.Transactions()),
	}
	resp = new(types.Response)
	resp.Data = info
	//// 打印区块信息
	//fmt.Println("Block Number:", blockInfo["number"])
	//fmt.Println("Parent Hash:", blockInfo["parentHash"])
	//fmt.Println("Timestamp:", blockInfo["timestamp"])
	//fmt.Println("Transactions:", blockInfo["transactions"])
	return
}
