package bo

type GetBlockLatestNReq struct {
	Limit int `json:"limit"`
}

type GetBlockLatestNResp struct {
	Num        uint64 `json:"block_num" example:"34449189"`
	Hash       string `json:"block_hash" example:"0x0000000"`
	Time       uint64 `json:"block_time" example:"1631534170"`
	ParentHash string `json:"parent_hash" example:"0x0000000"`
}

type GetBlockInfoResp struct {
	Num          uint64   `json:"block_num" example:"34449189"`
	Hash         string   `json:"block_hash" example:"0x0000000"`
	Time         uint64   `json:"block_time" example:"1631534170"`
	ParentHash   string   `json:"parent_hash" example:"0x0000000"`
	Transactions []string `json:"transactions" example:"[0x0000000, 0x0000001, 0x0000002]"`
}
