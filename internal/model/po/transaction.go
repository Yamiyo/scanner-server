package po

type Transaction struct {
	ID

	BlockNum uint64  `gorm:"<-:create;column:block_num;index:idx_block_num"`
	TxHash   string  `gorm:"<-:create;column:tx_hash"`
	From     string  `gorm:"<-:create;column:from"`
	To       *string `gorm:"<-:create;column:to"`
	Value    string  `gorm:"<-:create;column:value"`
	Nonce    uint64  `gorm:"<-:create;column:nonce"`
	Data     string  `gorm:"<-:create;column:data"`
	//Logs     []interface{} `gorm:"<-:create;column:logs"`

	CreatedAt
	UpdatedAt
	DeletedAt
}

type TransactionLog struct {
	ID

	BlockNum uint64 `gorm:"<-:create;column:block_num;index:idx_block_num"`
	TxHash   string `gorm:"<-:create;column:tx_hash;index:idx_block_num"`
	Index    uint64 `gorm:"<-:create;column:index"`
	Data     string `gorm:"<-:create;column:data"`
	CreatedAt
	UpdatedAt
	DeletedAt
}
