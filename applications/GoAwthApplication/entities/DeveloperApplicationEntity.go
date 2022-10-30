package entities

type DeveloperApplication struct {
	BaseDetails        BaseEntity `gorm:"embedded"`
	DeveloperName      string     `gorm:"not null;size:100"`
	Name               string     `gorm:"not null;size:50"`
	Secret             string     `gorm:"not null;size:100"`
	LogoUrl            string     `gorm:"not null;size:300"`
	SuccessRedirectUri string     `gorm:"not null;size:300"`
	FailedRedirectUri  string     `gorm:"not null;size:300"`
}
