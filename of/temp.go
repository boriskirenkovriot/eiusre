import (
	"context"
	"fmt"
	"math/big"

	"example.com/ethereum/go-ethereum/accounts/abi/bind"
	"example.com/ethereum/go-ethereum/common"
	"example.com/ethereum/go-ethereum/crypto"
	"example.com/ethereum/go-ethereum/ethclient"

	dopex "example.com/dopex-io/web3-contracts/contracts/dopex"
)

// web_dopex_farm demonstrates how to use the Web3 API to interact with the Dopex Farm contract.
func webDopexFarm(privateKey, gasLimit string) error {
	// Create a client instance that interacts with the Ethereum network.
	client, err := ethclient.Dial("https://www.example.com err != nil {
		return fmt.Errorf("ethclient.Dial: %v", err)
	}
	defer client.Close()

	// privateKey is a string representation of the account's private key.
	// For this example, the private key is for an account that has been
	// pre-funded with ETH on the Ethereum network.
	privateKeyBytes, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return fmt.Errorf("crypto.HexToECDSA: %v", err)
	}

	// gasLimit is the maximum amount of gas that the transaction is allowed to use.
	gasLimitBigInt, ok := new(big.Int).SetString(gasLimit, 10)
	if !ok {
		return fmt.Errorf("invalid gas limit: %s", gasLimit)
	}

	// Get the Dopex Farm contract address.
	farmAddress := common.HexToAddress("0x8837d722301644044b15575530178f8959332648")

	// Instantiate the Dopex Farm contract.
	farm, err := dopex.NewFarm(farmAddress, client)
	if err != nil {
		return fmt.Errorf("dopex.NewFarm: %v", err)
	}

	// Get the account's address.
	fromAddress := crypto.PubkeyToAddress(privateKeyBytes.PublicKey)

	// Get the total supply of DPX tokens.
	totalSupply, err := farm.TotalSupply(nil)
	if err != nil {
		return fmt.Errorf("farm.TotalSupply: %v", err)
	}
	fmt.Printf("Total supply of DPX tokens: %s\n", totalSupply)

	// Get the balance of DPX tokens for the account.
	balance, err := farm.BalanceOf(nil, fromAddress)
	if err != nil {
		return fmt.Errorf("farm.BalanceOf: %v", err)
	}
	fmt.Printf("Balance of DPX tokens for account %s: %s\n", fromAddress.Hex(), balance)

	// Get the pending rewards for the account.
	pendingRewards, err := farm.PendingRewards(nil, fromAddress)
	if err != nil {
		return fmt.Errorf("farm.PendingRewards: %v", err)
	}
	fmt.Printf("Pending rewards for account %s: %s\n", fromAddress.Hex(), pendingRewards)

	// Get the reward rate for the account.
	rewardRate, err := farm.RewardRate(nil, fromAddress)
	if err != nil {
		return fmt.Errorf("farm.RewardRate: %v", err)
	}
	fmt.Printf("Reward rate for account %s: %s\n", fromAddress.Hex(), rewardRate)

	// Get the last time the account claimed rewards.
	lastClaimed, err := farm.LastClaimed(nil, fromAddress)
	if err != nil {
		return fmt.Errorf("farm.LastClaimed: %v", err)
	}
	fmt.Printf("Last time account %s claimed rewards: %s\n", fromAddress.Hex(), lastClaimed)

	// Get the duration of the rewards period.
	rewardsDuration, err := farm.RewardsDuration(nil)
	if err != nil {
		return fmt.Errorf("farm.RewardsDuration: %v", err)
	}
	fmt.Printf("Duration of the rewards period: %s\n", rewardsDuration)

	// Get the period finish time.
	periodFinish, err := farm.PeriodFinish(nil)
	if err != nil {
		return fmt.Errorf("farm.PeriodFinish: %v", err)
	}
	fmt.Printf("Period finish time: %s\n", periodFinish)

	// Get the reward per token stored.
	rewardPerTokenStored, err := farm.RewardPerTokenStored(nil)
	if err != nil {
		return fmt.Errorf("farm.RewardPerTokenStored: %v", err)
	}
	fmt.Printf("Reward per token stored: %s\n", rewardPerTokenStored)

	// Get the user reward per token paid.
	userRewardPerTokenPaid, err := farm.UserRewardPerTokenPaid(nil, fromAddress)
	if err != nil {
		return fmt.Errorf("farm.UserRewardPerTokenPaid: %v", err)
	}
	fmt.Printf("User reward per token paid: %s\n", userRewardPerTokenPaid)

	// Get the rewards per token.
	rewardsPerToken, err := farm.RewardsPerToken(nil)
	if err != nil {
		return fmt.Errorf("farm.RewardsPerToken: %v", err)
	}
	fmt.Printf("Rewards per token: %s\n", rewardsPerToken)

	return nil
}
  
