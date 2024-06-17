package logic

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"

	"EthTest/demo/internal/svc"
	"EthTest/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountInfoLogic {
	return &AccountInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type AccountInfo struct {
	Nonce   uint64 `json:"nonce"`
	Balance string `json:"balance"`
}

func (l *AccountInfoLogic) AccountInfo(req *types.AccountRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// 使用rpc的方式
	var client *ethclient.Client
	client, err = ethclient.Dial(l.svcCtx.Config.EthClient.BaseUrl) //  配置的eth地址
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	// 要查询的以太坊账户地址
	address := common.HexToAddress(req.EthAddr)
	// 获取账户的 nonce
	var nonce uint64
	var balance *big.Int
	balance, err = client.BalanceAt(l.ctx, address, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	nonce, err = client.NonceAt(l.ctx, address, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	info := AccountInfo{
		Nonce:   nonce,
		Balance: balance.String(),
	}
	resp = new(types.Response)
	resp.Data = info

	return
}
