package twallet

// #include <TrustWalletCore/TWEthereumProto.h>
// #include <TrustWalletCore/TWBitcoinProto.h>
// #include <TrustWalletCore/TWAnySigner.h>
import "C"
import (
	"github.com/golang/protobuf/proto"
	"unsafe"
)

var signers = map[string]func(unsafe.Pointer) unsafe.Pointer{
	"BTC":  btcSigner,
	"BCH":  bchSigner,
	"DASH": dashSigner,
	"DGB":  dgbSigner,
	"DOGE": dogeSigner,
	"ETH":  ethSigner,
	"GRS":  grsSigner,
	"LTC":  ltcSigner,
	"TRX":  trxSigner,
	"ZEC":  zecSigner,
	"RVN":  rvnSigner,
}

func signTransaction(coin string, pb proto.Message) ([]byte, error) {
	out, err := proto.Marshal(pb)
	if err != nil {
		return []byte(""), err
	}
	twData := TWDataCreateWithGoBytes(out)
	signer := signers[coin]
	result := signer(twData)
	return TWDataGoBytes(unsafe.Pointer(result)), nil
}

func anySigner(twData unsafe.Pointer, coinType uint32) unsafe.Pointer {
	return C.TWAnySignerSign(twData, coinType)
}

func ethSigner(twData unsafe.Pointer) unsafe.Pointer {
	return anySigner(twData, twCoinTypes["ETH"])
}

func btcSigner(twData unsafe.Pointer) unsafe.Pointer {
	return anySigner(twData, twCoinTypes["BTC"])
}

func zecSigner(twData unsafe.Pointer) unsafe.Pointer {
	return anySigner(twData, twCoinTypes["ZEC"])
}

func grsSigner(twData unsafe.Pointer) unsafe.Pointer {
	return anySigner(twData, twCoinTypes["GRS"])
}

func ltcSigner(twData unsafe.Pointer) unsafe.Pointer {
	return anySigner(twData, twCoinTypes["LTC"])
}

func bchSigner(twData unsafe.Pointer) unsafe.Pointer {
	return anySigner(twData, twCoinTypes["BCH"])
}

func dashSigner(twData unsafe.Pointer) unsafe.Pointer {
	return anySigner(twData, twCoinTypes["DASH"])
}

func dogeSigner(twData unsafe.Pointer) unsafe.Pointer {
	return anySigner(twData, twCoinTypes["DOGE"])
}

func dgbSigner(twData unsafe.Pointer) unsafe.Pointer {
	return anySigner(twData, twCoinTypes["DGB"])
}

func trxSigner(twData unsafe.Pointer) unsafe.Pointer {
	return anySigner(twData, twCoinTypes["TRX"])
}

func rvnSigner(twData unsafe.Pointer) unsafe.Pointer {
	return anySigner(twData, twCoinTypes["RVN"])
}
