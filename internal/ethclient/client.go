package ethclient

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func Connect(host string) (*ethclient.Client, error) {
	conn, err := ethclient.Dial(host)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
