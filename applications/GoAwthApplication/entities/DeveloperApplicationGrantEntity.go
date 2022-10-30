package entities

type DeveloperApplicationGrant struct {
	BaseDetails    BaseEntity `gorm:"embedded"`
	PlayerUsername string     `gorm:"not null;size:100"`
	ApplicationId  string     `gorm:"not null;size:100"`
	ExpiresAt      string     `gorm:"not null;size:100"`
	Scope          string     `gorm:"not null"`
}
