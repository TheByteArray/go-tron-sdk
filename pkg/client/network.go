package client

import (
	"bytes"
	"fmt"

	"github.com/TheByteArray/go-tron-sdk/pkg/common"
	"github.com/TheByteArray/go-tron-sdk/pkg/proto/api"
	"github.com/TheByteArray/go-tron-sdk/pkg/proto/core"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

// ListNodes provides list of network nodes
func (g *GrpcClient) ListNodes() (*api.NodeList, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	nodeList, err := g.Client.ListNodes(ctx, new(api.EmptyMessage))
	if err != nil {
		zap.L().Error("List nodes", zap.Error(err))
	}
	return nodeList, nil
}

// GetNextMaintenanceTime get next epoch timestamp
func (g *GrpcClient) GetNextMaintenanceTime() (*api.NumberMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetNextMaintenanceTime(ctx,
		new(api.EmptyMessage))
}

// TotalTransaction return total transciton in network
func (g *GrpcClient) TotalTransaction() (*api.NumberMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.TotalTransaction(ctx,
		new(api.EmptyMessage))
}

// GetTransactionByID returns transaction details by ID
func (g *GrpcClient) GetTransactionByID(id string) (*core.Transaction, error) {
	transactionID := new(api.BytesMessage)
	var err error

	transactionID.Value, err = common.FromHex(id)
	if err != nil {
		return nil, fmt.Errorf("get transaction by id error: %v", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.GetTransactionById(ctx, transactionID)
	if err != nil {
		return nil, err
	}
	if size := proto.Size(tx); size > 0 {
		return tx, nil
	}
	return nil, fmt.Errorf("transaction info not found")
}

// GetTransactionInfoByID returns transaction receipt by ID
func (g *GrpcClient) GetTransactionInfoByID(id string) (*core.TransactionInfo, error) {
	transactionID := new(api.BytesMessage)
	var err error

	transactionID.Value, err = common.FromHex(id)
	if err != nil {
		return nil, fmt.Errorf("get transaction by id error: %v", err)
	}

	ctx, cancel := g.getContext()
	defer cancel()

	txi, err := g.Client.GetTransactionInfoById(ctx, transactionID)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(txi.Id, transactionID.Value) {
		return txi, nil
	}
	return nil, fmt.Errorf("transaction info not found")
}

// Broadcast broadcast TX
func (g *GrpcClient) Broadcast(tx *core.Transaction) (*api.Return, error) {
	ctx, cancel := g.getContext()
	defer cancel()
	result, err := g.Client.BroadcastTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	if !result.GetResult() {
		return result, fmt.Errorf("result error: %s", result.GetMessage())
	}
	if result.GetCode() != api.Return_SUCCESS {
		return result, fmt.Errorf("result error(%s): %s", result.GetCode(), result.GetMessage())
	}
	return result, nil
}

// GetNodeInfo current connection
func (g *GrpcClient) GetNodeInfo() (*core.NodeInfo, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetNodeInfo(ctx, new(api.EmptyMessage))
}

// GetEnergyPrices returns energy prices
func (g *GrpcClient) GetEnergyPrices() (*api.PricesResponseMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetEnergyPrices(ctx, new(api.EmptyMessage))
}

// GetBandwidthPrices returns bandwidth prices
func (g *GrpcClient) GetBandwidthPrices() (*api.PricesResponseMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetBandwidthPrices(ctx, new(api.EmptyMessage))
}

// GetMemoFee returns memo fee
func (g *GrpcClient) GetMemoFee() (*api.PricesResponseMessage, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.GetMemoFee(ctx, new(api.EmptyMessage))
}

func GRPCInsecure() grpc.DialOption {
	return grpc.WithTransportCredentials(insecure.NewCredentials())
}
