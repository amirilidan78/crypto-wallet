package twallet

// #cgo CFLAGS: -I/wallet-core/include
// #cgo LDFLAGS: -L/wallet-core/build -L/wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWString.h>
// #include <TrustWalletCore/TWData.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
// #include <TrustWalletCore/TWBitcoinScript.h>
import "C"
import (
	"encoding/hex"
)

var twCoinTypes = map[string]uint32{
	"AE":    C.TWCoinTypeAeternity,
	"AION":  C.TWCoinTypeAion,
	"ALGO":  C.TWCoinTypeAlgorand,
	"BNB":   C.TWCoinTypeBinance,
	"BTC":   C.TWCoinTypeBitcoin,
	"BCH":   C.TWCoinTypeBitcoinCash,
	"BTG":   C.TWCoinTypeBitcoinGold,
	"CLO":   C.TWCoinTypeCallisto,
	"ADA":   C.TWCoinTypeCardano,
	"ATOM":  C.TWCoinTypeCosmos,
	"DASH":  C.TWCoinTypeDash,
	"DCR":   C.TWCoinTypeDecred,
	"DGB":   C.TWCoinTypeDigiByte,
	"DOGE":  C.TWCoinTypeDogecoin,
	"EOS":   C.TWCoinTypeEOS,
	"ETH":   C.TWCoinTypeEthereum,
	"ETC":   C.TWCoinTypeEthereumClassic,
	"FIO":   C.TWCoinTypeFIO,
	"GO":    C.TWCoinTypeGoChain,
	"GRS":   C.TWCoinTypeGroestlcoin,
	"ICX":   C.TWCoinTypeICON,
	"IOTX":  C.TWCoinTypeIoTeX,
	"Kava":  C.TWCoinTypeKava,
	"KIN":   C.TWCoinTypeKin,
	"LTC":   C.TWCoinTypeLitecoin,
	"MONA":  C.TWCoinTypeMonacoin,
	"NAS":   C.TWCoinTypeNebulas,
	"NULS":  C.TWCoinTypeNULS,
	"NANO":  C.TWCoinTypeNano,
	"NEAR":  C.TWCoinTypeNEAR,
	"NIM":   C.TWCoinTypeNimiq,
	"ONT":   C.TWCoinTypeOntology,
	"POA":   C.TWCoinTypePOANetwork,
	"QTUM":  C.TWCoinTypeQtum,
	"XRP":   C.TWCoinTypeXRP,
	"SOL":   C.TWCoinTypeSolana,
	"XLM":   C.TWCoinTypeStellar,
	"TON":   C.TWCoinTypeTON,
	"XTZ":   C.TWCoinTypeTezos,
	"THETA": C.TWCoinTypeTheta,
	"TT":    C.TWCoinTypeThunderToken,
	"NEO":   C.TWCoinTypeNEO,
	"TOMO":  C.TWCoinTypeTomoChain,
	"TRX":   C.TWCoinTypeTron,
	"VET":   C.TWCoinTypeVeChain,
	"VIA":   C.TWCoinTypeViacoin,
	"WAN":   C.TWCoinTypeWanchain,
	"ZEC":   C.TWCoinTypeZcash,
	"XZC":   C.TWCoinTypeZcoin,
	"ZIL":   C.TWCoinTypeZilliqa,
	"ZEL":   C.TWCoinTypeZelcash,
	"RVN":   C.TWCoinTypeRavencoin,
	"WAVES": C.TWCoinTypeWaves,
	"LUNA":  C.TWCoinTypeTerra,
	"ONE":   C.TWCoinTypeHarmony,
	"KSM":   C.TWCoinTypeKusama,
	"IOU":   C.TWCoinTypePolkadot,
	"FIL":   C.TWCoinTypeFilecoin,
	"ERD":   C.TWCoinTypeElrond,
}

func hDWalletKey(passphrase string) string {
	cPassphrase := TWStringCreateWithGoString(passphrase)
	defer C.TWStringDelete(cPassphrase)
	wallet := C.TWHDWalletCreate(128, cPassphrase)
	defer C.TWHDWalletDelete(wallet)
	return TWStringGoString(C.TWHDWalletMnemonic(wallet))
}

func hdWalletAddressForCoin(seed string, passphrase string, coin string) string {
	cSeed := TWStringCreateWithGoString(seed)
	cPassphrase := TWStringCreateWithGoString(passphrase)
	defer C.TWStringDelete(cSeed)
	defer C.TWStringDelete(cPassphrase)
	wallet := C.TWHDWalletCreateWithMnemonic(cSeed, cPassphrase)
	defer C.TWHDWalletDelete(wallet)
	address := C.TWHDWalletGetAddressForCoin(wallet, twCoinTypes[coin])
	return TWStringGoString(address)
}

func hdWalletScriptBuildForAddress(address string, coin string) []byte {
	cAddress := TWStringCreateWithGoString(address)
	cCoin := twCoinTypes[coin]
	defer C.TWStringDelete(cAddress)
	//For version 2.4.1
	twScript := C.TWBitcoinScriptLockScriptForAddress(cAddress, cCoin)
	//twScript := C.TWBitcoinScriptBuildForAddress(cAddress, cCoin)
	twData := C.TWBitcoinScriptData(twScript)
	return TWDataGoBytes(twData)
}

func hdWalletPrivateKeyForCoin(seed string, passphrase string, coin string) string {
	cSeed := TWStringCreateWithGoString(seed)
	cPassphrase := TWStringCreateWithGoString(passphrase)
	defer C.TWStringDelete(cSeed)
	defer C.TWStringDelete(cPassphrase)
	wallet := C.TWHDWalletCreateWithMnemonic(cSeed, cPassphrase)
	defer C.TWHDWalletDelete(wallet)
	key := C.TWHDWalletGetKeyForCoin(wallet, twCoinTypes[coin])
	keyData := C.TWPrivateKeyData(key)
	keyHex := hex.EncodeToString(TWDataGoBytes(keyData))
	return keyHex
}
