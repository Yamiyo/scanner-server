package bo

type GetTransactionsReq struct {
}

type GetTransactionInfoResp struct {
	BlockNum uint64            `json:"block_num" example:"34449189" format:"uint64"`
	TxHash   string            `json:"tx_hash" example:"0x0000000"`
	From     string            `json:"from" example:"0x0000000"`
	To       *string           `json:"to" example:"0x0000000"`
	Value    string            `json:"value" example:"0x0000000"`
	Nonce    uint64            `json:"nonce" example:"0x0000000" format:"uint64"`
	Data     []byte            `json:"data" example:"0x0000000"`
	Logs     []*TransactionLog `json:"logs" example:"[{\"index\": 0, \"data\": \"0x0000000\"}]"`
}

type TransactionLog struct {
	Index uint   `json:"index"`
	Data  string `json:"data"`
}
