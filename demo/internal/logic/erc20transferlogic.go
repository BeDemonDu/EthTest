package logic

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	mtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"

	"EthTest/demo/internal/svc"
	"EthTest/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ERC20TransferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewERC20TransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ERC20TransferLogic {
	return &ERC20TransferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ERC20TransferLogic) ERC20Transfer(req *types.ERC20TransferRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	var client *ethclient.Client
	client, err = ethclient.Dial(l.svcCtx.Config.EthClient.BaseUrl) //  配置的eth地址
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
	//
	contractABI, err := abi.JSON(strings.NewReader(erc20baseAbi))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
		return
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	toAddress := common.HexToAddress(req.ToAddress)
	data, err := contractABI.Pack("transfer", toAddress, big.NewInt(req.Amount))
	if err != nil {

		return nil, err
	}

	// 获取发送方的 nonce
	var nonce uint64
	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return
	}
	var gasPrice *big.Int
	// 设置 gas 价格和 gas 限制
	gasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	gasLimit := uint64(21368) // 交易默认 gas 限制

	contractAddress := common.HexToAddress(req.ContractAddress)
	// 创建部署合约的交易
	deployTx := mtypes.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// 对交易进行签名
	var chainID *big.Int
	chainID, err = client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	var signedTx *mtypes.Transaction
	signedTx, err = mtypes.SignTx(deployTx, mtypes.NewEIP155Signer(chainID), privateKey)
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
	return
}

var erc20baseAbi = `
[

    {

        "constant": false,

        "inputs": [

            {

                "name": "to",

                "type": "address"

            },

            {

                "name": "value",

                "type": "uint256"

            }

        ],

        "name": "transfer",

        "outputs": [

            {

                "name": "",

                "type": "bool"

            }

        ],

        "payable": false,

        "stateMutability": "nonpayable",

        "type": "function"

    }

]`
