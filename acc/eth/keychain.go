package eth

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
		}
	}

	keychain, err = importWallet(privateKeyECDSA)
	if err != nil {
		return KeyChain{}, fmt.Errorf("error during import wallet: %v", err)
	}

	return keychain, err
}
