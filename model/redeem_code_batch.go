package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameRedeemCodeBatch = "redeem_code_batch"

// RedeemCodeBatch mapped from table <redeem_code_batch>
type RedeemCodeBatch struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`        // 批次名称
	Description string    `json:"description"`  // 批次描述
	UsageLimit  int32     `json:"usage_limit"`  // 可使用次数
	TotalCount  int32     `json:"total_count"`  // 生成数量
	UsedCount   int32     `json:"used_count"`   // 已使用数量
	StartedAt   time.Time `json:"started_at"`   // 开始时间
	EndedAt     time.Time `json:"ended_at"`     // 结束时间
	Status      int32     `json:"status"`       // 状态：0，默认；1，启用；2，停用；
	CreatorID   int64     `json:"creator_id"`   // 创建人 ID
	CreatorName string    `json:"creator_name"` // 创建人名称
	CreatedAt   time.Time `json:"created_at"`   // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`   // 更新时间
}

// TableName RedeemCodeBatch's table name
func (*RedeemCodeBatch) TableName() string {
	return TableNameRedeemCodeBatch
}

func GetRedeemCodeBatchByID(db *gorm.DB, id int64) (*RedeemCodeBatch, error) {
	var RedeemCodeBatch RedeemCodeBatch
	err := db.Where("id = ?", id).First(&RedeemCodeBatch).Error
	return &RedeemCodeBatch, err
}

func GetRedeemCodeBatchList(db *gorm.DB, page int32, pageSize int32) (list []*RedeemCodeBatch, count int64) {
	db.Model(&RedeemCodeBatch{}).Count(&count).Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Order("id desc").Find(&list)
	return
}
