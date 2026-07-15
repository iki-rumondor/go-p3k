package models

type SystemSetting struct {
	ID        uint   `gorm:"primaryKey"`
	Key       string `gorm:"unique;size:64;not_null"`
	Value     string `gorm:"size:255"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}
