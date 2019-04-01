/*
 * @Description:
 * @Author: Moqi
 * @Date: 2018-12-12 10:35:05
 * @Email: str@li.cm
 * @Github: https://github.com/strugglerx
 * @LastEditors: Moqi
 * @LastEditTime: 2018-12-12 10:35:07
 */

package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"sort"
	"strings"
)

func VerifySha1(token, timestamp, nonce string) string {
	verify := []string{token, timestamp, nonce}
	sort.Strings(verify)
	res := strings.Join(verify, "")
	//返回验证数据
	return Sha1String(res)
}

func Sha1String(value string) string {

	d := sha1.New()
	d.Write([]byte(value))
	r := d.Sum(nil)
	result := hex.EncodeToString(r)
	return result
}

func Md5String(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	r := m.Sum(nil)
	result := hex.EncodeToString(r)
	return result
}

//aes加密解密

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

const aeskey = "3dsa124dsf6gfg6s"

//aes加密解密接口
func CustomAesEncrypt(data string) string {
	Customcrypto, err := AesEncrypt([]byte(data), []byte(aeskey))
	if err != nil {
		return "-1"
	}
	return base64.StdEncoding.EncodeToString(Customcrypto)
}

func CustomAesDecrypt(data string) string {
	byteData, _ := base64.StdEncoding.DecodeString(data)
	CustomDecrypt, err := AesDecrypt(byteData, []byte(aeskey))
	if err != nil {
		return "-1"
	}
	return string(CustomDecrypt)
}
