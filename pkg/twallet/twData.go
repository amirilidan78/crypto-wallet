package twallet

// #cgo CFLAGS: -I/wallet-core/include
// #cgo LDFLAGS: -L/wallet-core/build -L/wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWString.h>
// #include <TrustWalletCore/TWData.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
import "C"
import (
	"unsafe"
)

// C.TWString -> Go string
func TWStringGoString(s unsafe.Pointer) string {
	return C.GoString(C.TWStringUTF8Bytes(s))
}

// Go string -> C.TWString
func TWStringCreateWithGoString(s string) unsafe.Pointer {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))
	str := C.TWStringCreateWithUTF8Bytes(cStr)
	return str
}

// C.TWData -> Go byte[]
func TWDataGoBytes(d unsafe.Pointer) []byte {
	cBytes := C.TWDataBytes(d)
	cSize := C.TWDataSize(d)
	return C.GoBytes(unsafe.Pointer(cBytes), C.int(cSize))
}

// Go byte[] -> C.TWData
func TWDataCreateWithGoBytes(d []byte) unsafe.Pointer {
	cBytes := C.CBytes(d)
	defer C.free(unsafe.Pointer(cBytes))
	data := C.TWDataCreateWithBytes((*C.uchar)(cBytes), C.ulong(len(d)))
	return data
}
