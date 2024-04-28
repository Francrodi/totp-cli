package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
)

const TimeStep = 30
const CodeLength = 6

func Run(secret string, time int64) {
	var counter int64 = time / TimeStep
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Println("Error decoding secret!", err)
		return
	}
	hmacResult := getHmacHash(key, counter)
	fmt.Printf("Code: %v\n", truncateHmacResult(hmacResult))
}

func getHmacHash(key []byte, counter int64) []byte {
	hmacHash := hmac.New(sha1.New, key)
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(counter))
	hmacHash.Write(data)
	hmacResult := hmacHash.Sum(nil)
	return hmacResult
}

func truncateHmacResult(hmacResult []byte) int {
	offset := hmacResult[len(hmacResult)-1] & 0x0f
	resultBytes := hmacResult[offset : offset+4]
	resultBytes[0] = resultBytes[0] & 0x7f

	code := int(binary.BigEndian.Uint32(resultBytes)) % int(math.Pow10(CodeLength))
	return code
}
