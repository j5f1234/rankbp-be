package model

type HeroBase struct {
	BaseModel

	HeroId    int    `gorm:"column:hero_id;not null;comment:'英雄id'" json:"hero_id"`
	Name      string `gorm:"column:name;not null;comment:'英雄名'" json:"name"`
	Position1 string `gorm:"column:position1;not null;comment:'位置1'" json:"position1"`
	Position2 string `gorm:"column:position2;not null;comment:'位置2'" json:"position2"`
}

func (HeroBase) TableNAme() string {
	return "herobases"
}
