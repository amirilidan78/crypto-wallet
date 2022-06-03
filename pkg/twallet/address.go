package twallet

// #include <TrustWalletCore/TWAnyAddress.h>
import "C"

func isAddressValid(address string, coin string) bool {
	coinCode := twCoinTypes[coin]
	result := C.TWAnyAddressIsValid(TWStringCreateWithGoString(address), coinCode)
	return bool(result)
}
