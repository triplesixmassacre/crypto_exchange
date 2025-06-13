package eth

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type KeyChain struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
	Address    common.Address
}

func importWallet(privateKeyECDSA *ecdsa.PrivateKey) (keychain KeyChain, err error) {
	pubKey := privateKeyECDSA.Public()
	fmt.Println("Public Key:", pubKey)

	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return KeyChain{}, fmt.Errorf("Failed to generate public key: %T", err)
	}

	fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
	fmt.Println("Address:", fromAddress)

	keychain = KeyChain{
		PrivateKey: *privateKeyECDSA,
		PublicKey:  *pubKeyECDSA,
		Address:    fromAddress,
	}
	return keychain, nil
}

func ImportOrGenerateWallet(privateKeyHex *string) (keychain KeyChain, err error) {
	var privateKeyECDSA *ecdsa.PrivateKey

	if privateKeyHex == nil {
		privateKeyECDSA, err = crypto.GenerateKey()
		log.Printf("New Private Key: %v", privateKeyECDSA)
		if err != nil {
			log.Fatalf("Error during create key: %v", err)
		}
	} else {
		privateKeyECDSA, err = crypto.HexToECDSA(*privateKeyHex)
		if err != nil {
			fmt.Errorf("Failed to load private key: %T", err)
		}
	}

	keychain, err = importWallet(privateKeyECDSA)
	if err != nil {
		log.Fatalf("Error during import wallet: %v", err)
	}
	fmt.Println(keychain)

	return keychain, err
}

func double_main() {
	// Generate a new wallet
	privateKey := "d34157f1b67603f397843b7aa4142e63aaac8ca9cf423b31ab1f4ae1e50dda17"

	keychain, err := ImportOrGenerateWallet(&privateKey)
	if err != nil {
		log.Fatalf("Error generating wallet: %v", err)
	}

	log.Printf("Added wallet address: %s", keychain.Address.Hex())

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/05dbddbaca1a47f0bc7421cb9b54a13b")
	defer client.Close()

	// toAddress := "0xA964B46Cd5C7a6E8C22547518cD697a84C8C2c0e"

	// txHash, err := eth.CreateTx(keychain, client, &toAddress, 0.005)
	// if err != nil {
	// 	log.Fatalf("Error during create tx: %v", err)
	// }
	// log.Println(txHash)

	balance, err := Balance(client, keychain)
	if err != nil {
		log.Fatalf("Error generating wallet: %v", err)
	}
	log.Printf("Eth balance: %v", balance.String())
}
