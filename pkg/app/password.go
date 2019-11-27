package app

import (
	"encoding/hex"
	"hash"
)

type Alger func() hash.Hash
type Crypt struct {
	Alg Alger
}

func (self *Crypt) Proc(password string) string {
	m := self.Alg()
	m.Write([]byte(password))
	st := m.Sum(nil)
	return hex.EncodeToString(st)
}

func (self *Crypt) Check(password, cipher string) bool {

	return self.Proc(password) == cipher
}
