package model

type HeroPool struct {
	BaseModel

	Position string `gorm:"column:position;not null;comment:'分路'" json:"position"`
	Name     string `gorm:"column:name;not null;comment:'英雄名'" json:"name"`
	Skilled  int    `gorm:"column:skilled;not null;comment:'熟练度，0-低，1-高';type:tinyint(2)" json:"skilled"`
}

func (HeroPool) TableNAme() string {
	return "heropools"
}
