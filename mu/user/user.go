package user

import (
	ss "github.com/shadowsocks/shadowsocks-go/shadowsocks"
)

var (
	client = NewClient()
)

func NewClient() Client {
	var client Client
	return client
}

func GetClient() Client {
	return client
}

func SetClient(c Client) {
	client = c
}

type Client interface {
	GetUsers() ([]User, error)
}

type User interface {
	GetPort() int
	GetPasswd() string
	GetMethod() string
	IsEnable() bool
	GetCipher() (*ss.Cipher, error)
	UpdateTraffic(storageSize int) error
	GetUserInfo() UserInfo
}

type UserInfo struct {
	Passwd string
	Port   int
	Method string
}
