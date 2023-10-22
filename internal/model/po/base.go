package po

import (
	"gorm.io/gorm"
	"time"
)

type ID struct {
	ID uint64 `gorm:"<-:create;column:id;primary_key;NOT NULL;comment:表格不重複主鍵"`
}

type CreatedAt struct {
	CreatedAt *time.Time `gorm:"<-:create;column:created_at;type:DATETIME(6);index:idx_order,sort:desc;NOT NULL;default:CURRENT_TIMESTAMP(6);comment:資料產生時間點"`
}

type UpdatedAt struct {
	UpdatedAt *time.Time `gorm:"column:updated_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);comment:資料最後修改時點"`
}

type DeletedAt struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:DATETIME(6);comment:資料軟刪除時間點"`
}
