package bo

type GetTransactionsReq struct{}

type GetTransactionInfoResp struct {
	BlockNum uint64            `json:"block_num" example:"34449189" format:"uint64"`
	TxHash   string            `json:"tx_hash" example:"0x0000000"`
	From     string            `json:"from" example:"0x0000001"`
	To       *string           `json:"to" example:"0x0000002"`
	Value    string            `json:"value" example:"0x0000003"`
	Nonce    string            `json:"nonce" example:"0x0000004"`
	Data     []byte            `json:"data"`
	Logs     []*TransactionLog `json:"logs"`
}

type TransactionLog struct {
	Index uint   `json:"index" example:"0"`
	Data  string `json:"data" example:"0x0000006"`
}
