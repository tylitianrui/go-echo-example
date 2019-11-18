package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
)

type Encrypter interface {
	Encode(b []byte) string
	Compare(password, cipher string) bool
}

type MD5Encrypt struct {
	instance hash.Hash
}

func (self *MD5Encrypt) Encode(password string) string {
	m := self.instance.Sum([]byte(password))
	return hex.EncodeToString(m[:])
}

func (self *MD5Encrypt) Compare(password, cipher string) bool {
	return self.Encode(password) == cipher

}

func NewMD5() *MD5Encrypt {
	return &MD5Encrypt{
		instance: md5.New(),
	}
}
