package eth

import (
	"crypto/ecdsa"
	"fmt"
	"log"
<<<<<<< HEAD
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
=======

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
>>>>>>> 818eb69afb0aa3a363e0a4e046f986d202e43f22
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

<<<<<<< HEAD
func preparePrivateKey(s string) (preparedPrivateKey string, err error) {
	// Проверка корректности введенного PrivateKey
	// Функция убирает все спец и запрещенные символы
	// PrivateKey 4битный, включает в себя 0-9 и a-f, остальные буквы не могут
	// находиться в private key впринципе
	lowercase := strings.ToLower(s)

	cleaned := strings.TrimSpace(lowercase)
	cleaned = regexp.MustCompile(`[\s]+`).ReplaceAllString(cleaned, "")

	// Условие ниже описывает случай, когда строка после удаления всех пробелов
	// оказалось пустой, пустая строка в параметре ключа означает генерацию
	// нового приватного ключа, поэтому возвращается пустая строка без ошибок
	if cleaned == "" {
		return "", nil
	}

	if !regexp.MustCompile(`^[a-f0-9]+$`).MatchString(cleaned) {
		return "", fmt.Errorf("error: string contains invalid characters (only a-f and 0-9 allowed)")
	}

	if len(cleaned) != 64 {
		return "", fmt.Errorf("incorrect input len")
	}

	return s, nil
}

func ImportOrGenerateWallet(privateKeyHex string) (keychain KeyChain, err error) {
	var privateKeyECDSA *ecdsa.PrivateKey

	preparePrivateKeyHex, err := preparePrivateKey(privateKeyHex)
	if err != nil {
		log.Fatalf("Error during prepare key: %v", err)
	}

	if preparePrivateKeyHex == "" {
		privateKeyECDSA, err = crypto.GenerateKey()
		// log.Printf("New Private Key: %v", privateKeyECDSA)
		if err != nil {
			return KeyChain{}, fmt.Errorf("error during create key: %v", err)
		}
	} else {
		if err != nil {
			return KeyChain{}, fmt.Errorf("failed prepare private key: %T", err)
		}

		privateKeyECDSA, err = crypto.HexToECDSA(preparePrivateKeyHex)
		if err != nil {
			return KeyChain{}, fmt.Errorf("failed to load private key: %T", err)
=======
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
>>>>>>> 818eb69afb0aa3a363e0a4e046f986d202e43f22
		}
	}

	keychain, err = importWallet(privateKeyECDSA)
	if err != nil {
<<<<<<< HEAD
		return KeyChain{}, fmt.Errorf("error during import wallet: %v", err)
	}

	return keychain, err
}
=======
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
>>>>>>> 818eb69afb0aa3a363e0a4e046f986d202e43f22
