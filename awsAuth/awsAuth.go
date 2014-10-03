package awsAuth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func getSignature(data string) string {
	return ComputeHmac256(data, SecretKey)
}

func AuthorizationString(data string) string {
	signature := getSignature(data)

	return "AWS3-HTTPS AWSAccessKeyId=" + AccessKey + ",Algorithm=HmacSHA256,Signature=" + signature
}
