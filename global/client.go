package global

import "github.com/ethereum/go-ethereum/ethclient"

var (
	Client *ethclient.Client
)

func ConnectToNode(grpc string) error {
	var err error
	Client, err = ethclient.Dial(grpc)
	if err != nil {
		return err
	}
	return nil
}
