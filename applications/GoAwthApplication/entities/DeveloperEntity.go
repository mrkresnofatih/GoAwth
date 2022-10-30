package entities

type Developer struct {
	BaseDetails           BaseEntity             `gorm:"embedded"`
	DeveloperName         string                 `gorm:"not null;uniqueIndex;size:50"`
	Password              string                 `gorm:"not null;size:300"`
	DeveloperApplications []DeveloperApplication `gorm:"foreignKey:DeveloperName;references:DeveloperName"`
}
