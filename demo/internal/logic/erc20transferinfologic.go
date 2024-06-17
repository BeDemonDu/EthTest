package logic

import (
	"EthTest/demo/internal/svc"
	"EthTest/demo/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	mtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/zeromicro/go-zero/rest/httpx"
	"log"
	"math/big"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type ERC20TransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

type ERC20TransferInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewERC20TransferInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ERC20TransferInfoLogic {
	return &ERC20TransferInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var abiValue = `[{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`

type T struct {
	Inputs []struct {
		Components []struct {
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
		} `json:"components,omitempty"`
		Indexed      bool   `json:"indexed,omitempty"`
		InternalType string `json:"internalType"`
		Name         string `json:"name"`
		Type         string `json:"type"`
	} `json:"inputs,omitempty"`
	StateMutability string `json:"stateMutability,omitempty"`
	Type            string `json:"type"`
	Anonymous       bool   `json:"anonymous,omitempty"`
	Name            string `json:"name,omitempty"`
	Outputs         []struct {
		InternalType string `json:"internalType"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Components   []struct {
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
		} `json:"components,omitempty"`
	} `json:"outputs,omitempty"`
}

func (l *ERC20TransferInfoLogic) ERC20TransferInfo(req *types.ERC20TransferInfoRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	var client *ethclient.Client
	client, err = ethclient.Dial(l.svcCtx.Config.EthClient.BaseUrl) //  配置的eth地址
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	erc20ABI, err := abi.JSON(strings.NewReader(abiValue))

	if err != nil {
		log.Fatal(err)
		return

	}

	contractAddress := common.HexToAddress(req.ContractAddress)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(req.StartBlock),
		ToBlock:   big.NewInt(req.EndBlock),
		Addresses: []common.Address{contractAddress},
	}
	var logs []mtypes.Log
	logs, err = client.FilterLogs(l.ctx, query)

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(len(logs), "tetst")
	if len(logs) == 0 {
		fmt.Println("etst")
		return
	}

	var events []ERC20TransferEvent

	for _, vLog := range logs {

		if len(vLog.Data) < 32 {
			fmt.Println("Data length insufficient: got %d, want 32 or more", len(vLog.Data))
		}
		var event ERC20TransferEvent
		err3 := erc20ABI.UnpackIntoInterface(&event, "Transfer", vLog.Data)

		if err3 != nil {
			fmt.Println(err3.Error(), l.ctx)
			log.Fatal(err3)
			return nil, err3
		}
		event.From = common.HexToAddress(vLog.Topics[1].Hex())
		event.To = common.HexToAddress(vLog.Topics[2].Hex())

		events = append(events, event)
	}

	resp = new(types.Response)
	resp.Data = events

	return
}
