package po

type Transaction struct {
	ID

	BlockNum uint64  `gorm:"<-:create;column:block_num;index:idx_block_num"`
	TxHash   string  `gorm:"<-:create;column:tx_hash"`
	From     string  `gorm:"<-:create;column:from"`
	To       *string `gorm:"<-:create;column:to;NULL"`
	Value    string  `gorm:"<-:create;column:value"`
	Nonce    uint64  `gorm:"<-:create;column:nonce"`
	Data     []byte  `gorm:"<-:create;column:data;type:blob"`

	CreatedAt
	UpdatedAt
	DeletedAt
}

type TransactionLog struct {
	ID

	BlockNum uint64 `gorm:"<-:create;column:block_num;index:idx_block_num"`
	TxHash   string `gorm:"<-:create;column:tx_hash;index:idx_txn_hash"`
	Index    uint   `gorm:"<-:create;column:index"`
	Log      []byte `gorm:"<-:create;column:logs;type:json"`

	CreatedAt
	UpdatedAt
	DeletedAt
}
