package po

type Block struct {
	ID

	Num        uint64 `gorm:"<-:create;column:num;index:unique, idx_num"`
	Hash       string `gorm:"<-:create;column:hash"`
	Time       uint64 `gorm:"<-:create;column:time"`
	ParentHash string `gorm:"<-:create;column:parent_hash"`

	CreatedAt
	UpdatedAt
	DeletedAt
}
