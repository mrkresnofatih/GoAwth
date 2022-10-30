package entities

type Player struct {
	BaseDetails BaseEntity `gorm:"embedded"`
	Username    string     `gorm:"not null,uniqueIndex;size:50"`
	FullName    string     `gorm:"not null;size:100"`
	ImageUrl    string     `gorm:"not null;size:300"`
	Password    string     `gorm:"not null;size:300"`
}
