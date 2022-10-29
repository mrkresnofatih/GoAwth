package entities

type BaseEntity struct {
	Id        string `gorm:"primaryKey"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"`
}
