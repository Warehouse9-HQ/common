package blockchain

type Block struct {
	// Basics:
	Timestamp    int64  `json:"timestamp"`     //Timestamp to control timing
	MinerAddress string `json:"miner_address"` // Address of node who mined it

	BlockNumber int64 `json:"block_number"`

	Hash              string `json:"hash"`                // Hash of block
	PreviousBlockHash string `json:"previous_block_hash"` // Previous block hash
	CutSignature      string `json:"cut_signature"`       // Only timestamp + miner_address + hash + previous_block_hash signed. Will be used to quickly verify downloaded blocks
	FullSignature     string `json:"full_signature"`      // Full block signature

	Transactions []Transaction `json:"transactions"` // All transactions

	RewardDistribution map[string]int64 `json:"reward_distribution"` // Reward distribution scheme
}
