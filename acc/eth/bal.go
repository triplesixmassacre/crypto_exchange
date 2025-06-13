package eth

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Balance(client *ethclient.Client, keychain KeyChain) (balance *big.Int, err error) {
	balance, err = client.BalanceAt(context.Background(), keychain.Address, nil)
	if err != nil {
		log.Fatalf("Error during get balace at: %s; error: %v", keychain.Address, err)
	}
	return balance, nil
}
