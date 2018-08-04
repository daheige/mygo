// php版本写法: $s = hash_hmac('sha256', 'Message', 'secret', true);
// echo base64_encode($s);

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	fmt.Println(ComputeHmac256("Message", "secret"))
}
