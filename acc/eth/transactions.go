package eth

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CreateTx(keychain KeyChain, conn *ethclient.Client, toAddressString *string, value float64) (txHash string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	nonce, err := conn.PendingNonceAt(ctx, keychain.Address)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	fmt.Println(value * 1_000_000_000_000_000_000)
	wei_value := big.NewInt(1_000_000_000_000_000_000)
	gasLimit := uint64(21000)

	gasTipCap, err := conn.SuggestGasTipCap(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get gas tip cap: %v", err)
	}
	gasFeeCap, err := conn.SuggestGasPrice(ctx) // Приблизительно BaseFee + GasTipCap
	if err != nil {
		return "", fmt.Errorf("failed to get gas fee cap: %v", err)
	}

	chainID, err := conn.NetworkID(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %v", err)
	}

	toAddress := common.HexToAddress(*toAddressString)

	txData := types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     wei_value,
		Data:      nil,
	}

	signer := types.NewLondonSigner(chainID)
	signedTx, err := types.SignNewTx(&keychain.PrivateKey, signer, &txData)
	if err != nil {
		return "", fmt.Errorf("failed to sign tx: %v", err)
	}

	err = conn.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send tx: %v", err)
	}

	return signedTx.Hash().Hex(), nil
}
