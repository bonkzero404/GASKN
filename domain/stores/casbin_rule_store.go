package stores

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:128;index"`
	V0    string `gorm:"size:128;index"`
	V1    string `gorm:"size:128;index"`
	V2    string `gorm:"size:128;index"`
	V3    string `gorm:"size:128;index"`
	V4    string `gorm:"size:128;index"`
	V5    string `gorm:"size:128;index"`
}
