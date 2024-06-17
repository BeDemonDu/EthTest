package logic

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	mtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"

	"EthTest/demo/internal/svc"
	"EthTest/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EthTradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEthTradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EthTradeLogic {
	return &EthTradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EthTradeLogic) EthTrade(req *types.TradeRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	var client *ethclient.Client
	client, err = ethclient.Dial(l.svcCtx.Config.EthClient.BaseUrl) // 8545 是默认的以太坊节点 RPC 端口
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	// 发送方地址和私钥（用于签名交易）
	fromPrivateKeyHex := req.PrivateKey
	privateKey, err := crypto.HexToECDSA(fromPrivateKeyHex)
	if err != nil {
		log.Fatal(err)
		return
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	// 接收方地址
	toAddress := common.HexToAddress(req.TargetAddr)
	// 转账金额
	amount := big.NewInt(req.Amount) // 1 ETH
	fmt.Println(fromAddress)
	fmt.Println(toAddress)
	fmt.Println(privateKey)

	// 获取发送方的 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 设置 gas 价格和 gas 限制
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	gasLimit := uint64(21000) // 交易默认 gas 限制

	// 创建转账交易
	tx := mtypes.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	// 对交易进行签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	signedTx, err := mtypes.SignTx(tx, mtypes.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 将签名后的交易发送到以太坊网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
		return
	}
	resp = new(types.Response)
	resp.Data = signedTx.Hash().Hex() // 返回对应的hash
	return
}
