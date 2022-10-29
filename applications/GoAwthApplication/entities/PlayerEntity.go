package entities

type PlayerEntity struct {
	BaseDetails BaseEntity `gorm:"embedded"`
	Username    string     `gorm:"unique,size:50"`
	FullName    string     `gorm:"not null,size:100"`
	ImageUrl    string     `gorm:"not null,size:250"`
	Password    string     `gorm:"not null,size:250"`
}
