package soft

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

const (
	Base    = "Base"
	Alpha   = "Alpha"   //初级版本 内测版。
	Beta    = "Beta"    //公测版
	RC      = "RC"      //Release Candidate 候选版本
	Release = "Release" //最终版本
	Free    = "Free"    //免费版。
	Full    = "Full"    //完全版。
)

type Version struct {
	Version   string `json:"tag"`
	Log       string `json:"log"`
	Status    string `json:"status"`
	Hash      string `json:"hash"`
	GitHash   string `json:"git_hash"`
	CreatedAt string `json:"created_at"`
}

func Md5FileStr() (string, error) {
	src, err := os.Open(os.Args[0])
	if err != nil {
		return "", err
	}
	instance := md5.New()
	_, err = io.Copy(instance, src)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(instance.Sum(nil)), nil
}
