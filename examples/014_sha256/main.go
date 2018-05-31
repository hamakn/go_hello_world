package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte(""))
	b := h.Sum(nil)

	s := base64.URLEncoding.EncodeToString(b)
	fmt.Println(s)
	// => 47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU=
	// ruby -r "digest" -e "p Digest::SHA256.base64digest(\"\")" に一致

	fmt.Println(hex.EncodeToString(b))
	// => e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
	// ruby -r "digest" -e "p Digest::SHA256.hexdigest(\"\")" に一致
}
