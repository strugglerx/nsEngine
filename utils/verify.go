package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

func VerifySha1(token,timestamp,nonce string) string {
	verify :=[]string{token,timestamp,nonce}
	sort.Strings(verify)
	res:=strings.Join(verify,"")
	return res
}

func Sha1String(value string) string {

	d:=sha1.New()
	d.Write([]byte(value))
	r:=d.Sum(nil)
	result:=hex.EncodeToString(r)
	return result
}

func Md5String(value string)string  {
	m:=md5.New()
	m.Write([]byte(value))
	r:=m.Sum(nil)
	result:=hex.EncodeToString(r)
	return result
}
