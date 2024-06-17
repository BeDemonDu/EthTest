package logic

import (
	"EthTest/demo/internal/svc"
	"EthTest/demo/internal/types"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type EthTradeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEthTradeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EthTradeInfoLogic {
	return &EthTradeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EthTradeInfoLogic) EthTradeInfo(req *types.TradeInfoRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	var client *rpc.Client
	client, err = rpc.Dial(l.svcCtx.Config.EthClient.BaseUrl) // 8545 是默认的以太坊节点 RPC 端口
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	// 要查询的交易哈希
	txHash := common.HexToHash(req.TxHash)

	// 获取交易信息
	var txInfo map[string]interface{}
	err = client.CallContext(context.Background(), &txInfo, "eth_getTransactionByHash", txHash)
	if err != nil {
		log.Fatal(err)
		return
	}
	resp = new(types.Response)
	resp.Data = txInfo

	// 打印交易信息
	//fmt.Println("Transaction Hash:", txInfo["hash"])
	//fmt.Println("Block Number:", txInfo["blockNumber"])
	//fmt.Println("From:", txInfo["from"])
	//fmt.Println("To:", txInfo["to"])
	//fmt.Println("Value:", txInfo["value"])
	//fmt.Println("Gas:", txInfo["gas"])
	//fmt.Println("Gas Price:", txInfo["gasPrice"])
	//fmt.Println("Input Data:", txInfo["input"])
	return
}
