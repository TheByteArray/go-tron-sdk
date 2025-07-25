package keys

import (
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/TheByteArray/go-tron-sdk/pkg/address"
	"github.com/TheByteArray/go-tron-sdk/pkg/common"
	"github.com/TheByteArray/go-tron-sdk/pkg/keystore"

	// "github.com/ethereum/go-ethereum/crypto"

	homedir "github.com/mitchellh/go-homedir"
)

func checkAndMakeKeyDirIfNeeded() string {
	userDir, _ := homedir.Dir()
	tronCTLDir := path.Join(userDir, ".tronctl", "keystore")
	if _, err := os.Stat(tronCTLDir); os.IsNotExist(err) {
		// Double check with Leo what is right file persmission
		err := os.Mkdir(tronCTLDir, 0700)
		if err != nil {
			fmt.Printf("create keystore dir error: %v\n", err)
			return ""
		}
	}

	return tronCTLDir
}

func ListKeys(keystoreDir string) {
	tronCTLDir := checkAndMakeKeyDirIfNeeded()
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(tronCTLDir, scryptN, scryptP)
	// keystore.KeyStore
	allAccounts := ks.Accounts()
	fmt.Printf("Tron Address:%s File URL:\n", strings.Repeat(" ", address.AddressLengthBase58))
	for _, account := range allAccounts {
		fmt.Printf("%s\t\t %s\n", account.Address, account.URL)
	}
}

func AddNewKey(password string) {
	tronCTLDir := checkAndMakeKeyDirIfNeeded()
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(tronCTLDir, scryptN, scryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		fmt.Printf("new account error: %v\n", err)
	}
	fmt.Printf("account: %s\n", account.Address)
	fmt.Printf("URL: %s\n", account.URL)
}

func GenerateKey() (*btcec.PrivateKey, error) {
	return btcec.NewPrivateKey()
}

func GetPrivateKeyFromHex(privateKeyHex string) (*btcec.PrivateKey, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key hex: %w", err)
	}

	return GetPrivateKeyFromBytes(privateKeyBytes)
}

func GetPrivateKeyFromBytes(privateKeyBytes []byte) (*btcec.PrivateKey, error) {
	if len(privateKeyBytes) != 32 {
		return nil, fmt.Errorf("invalid private key length: %d", len(privateKeyBytes))
	}

	if len(privateKeyBytes) != common.Secp256k1PrivateKeyBytesLength {
		return nil, common.ErrBadKeyLength
	}

	// btcec.PrivKeyFromBytes only returns a secret key and public key
	private, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	if private == nil {
		return nil, fmt.Errorf("failed to create private key from bytes")
	}

	return private, nil
}
