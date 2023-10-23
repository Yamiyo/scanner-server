package bo

type GetBlockLatestNReq struct {
	Limit int `json:"limit"`
}

type GetBlockLatestNResp struct {
	Num        uint64 `json:"block_num"`
	Hash       string `json:"block_hash"`
	Time       uint64 `json:"block_time"`
	ParentHash string `json:"parent_hash"`
}

type GetBlockInfoResp struct {
	Num          uint64   `json:"block_num"`
	Hash         string   `json:"block_hash"`
	Time         uint64   `json:"block_time"`
	ParentHash   string   `json:"parent_hash"`
	Transactions []string `json:"transactions"`
}
