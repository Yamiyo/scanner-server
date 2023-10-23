package bo

type GetTransactionsReq struct {
}

type GetTransactionInfoResp struct {
	BlockNum uint64            `json:"block_num"`
	TxHash   string            `json:"tx_hash"`
	From     string            `json:"from"`
	To       *string           `json:"to"`
	Value    string            `json:"value"`
	Nonce    uint64            `json:"nonce"`
	Data     []byte            `json:"data"`
	Logs     []*TransactionLog `json:"logs"`
}

type TransactionLog struct {
	Index uint   `json:"index"`
	Data  string `json:"data"`
}
