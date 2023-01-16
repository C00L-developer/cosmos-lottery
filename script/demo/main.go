package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"lottery/app"
	lotterytypes "lottery/x/lottery/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"
)

type ConditionFunc func() (done bool, err error)

var DefaultNodeHome string

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".lottery")
}

func main() {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",    // your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		panic(err)
	}
	defer grpcConn.Close()

	for count := 0; count < 100; count++ {
		for i := 1; i <= 20; i++ {
			go func() {
				if err := exec.Command("lotteryd", "tx", "lottery", "add-bet", fmt.Sprintf("%dtoken", i+5), "suffix", "--from", fmt.Sprintf("client%d", i), "--yes").Run(); err != nil {
					panic(err)
				}
			}()
			time.Sleep(10 * time.Millisecond)
		}
		if err := Poll(200*time.Millisecond, 5*time.Second, func() (bool, error) {
			lottery, err := getLottery(grpcConn)
			if lottery == nil {
				return false, nil
			}
			return lottery.Winner != "", err
		}); err != nil {
			break
		}
		for i := 1; i <= 20; i++ {
			go func() {
				if err := exec.Command("lotteryd", "tx", "lottery", "reveal-bet", "suffix", "--from", fmt.Sprintf("client%d", i), "--yes").Run(); err != nil {
					panic(err)
				}
			}()
			time.Sleep(10 * time.Millisecond)
		}
		if err := Poll(200*time.Millisecond, 5*time.Second, func() (bool, error) {
			lottery, err := getLottery(grpcConn)
			if lottery == nil {
				return false, nil
			}
			return lottery.Winner == "", err
		}); err != nil {
			break
		}
		fmt.Printf("Lottery %d is finished, Balances:\n", count+1)
		displayBalances(grpcConn)
	}

	lots, err := getAllLotteries(grpcConn)
	if err != nil {
		panic(err)
	}
	for _, lot := range lots {
		fmt.Printf("Lottery: %+v\n", lot)
	}
	displayBalances(grpcConn)
}

// Poll retries the given condition with the given interval until it succeeds
// or the given deadline expires.
func Poll(interval, deadline time.Duration, condition ConditionFunc) error {
	timeout := time.After(deadline)
	tick := time.NewTicker(interval)

	for {
		select {
		case <-timeout:
			return fmt.Errorf("timeout has been reached")
		case <-tick.C:
			ok, err := condition()
			if err != nil {
				return err
			}
			if ok {
				return nil
			}
		}
	}
}

func displayBalances(conn *grpc.ClientConn) {
	encodingConfig := app.MakeEncodingConfig()
	cliContext := client.Context{}.WithKeyringDir(DefaultNodeHome).WithCodec(encodingConfig.Marshaler)
	kr, _ := client.NewKeyringFromBackend(cliContext, "test")
	for i := 1; i <= 20; i++ {
		acc, _, _, _ := client.GetFromFields(cliContext, kr, fmt.Sprintf("client%d", i))
		bal, err := getBalance(conn, acc.String())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Client%d Account: %s Balance: %v\n", i, acc, bal)
	}
}

func getLottery(conn *grpc.ClientConn) (*lotterytypes.Lottery, error) {
	client := lotterytypes.NewQueryClient(conn)
	lotteryRes, err := client.CurrentLottery(
		context.Background(),
		&lotterytypes.QueryCurrentLotteryRequest{},
	)
	if err != nil {
		return nil, err
	}

	return &lotteryRes.Lottery, nil
}

func getAllLotteries(conn *grpc.ClientConn) ([]lotterytypes.Lottery, error) {
	client := lotterytypes.NewQueryClient(conn)
	lotteryRes, err := client.LotteryAll(
		context.Background(),
		&lotterytypes.QueryAllLotteryRequest{},
	)
	if err != nil {
		return nil, err
	}

	return lotteryRes.Lottery, nil
}

func getBalance(conn *grpc.ClientConn, addr string) (*sdk.Coin, error) {
	client := banktypes.NewQueryClient(conn)
	bankRes, err := client.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{Address: addr, Denom: "token"},
	)
	if err != nil {
		return nil, err
	}

	return bankRes.GetBalance(), nil
}
