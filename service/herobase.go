package service

import (
	"errors"
	"rankbp-be/common"
	"rankbp-be/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HeroBaseService struct {
}

// 更新英雄基础数据
func (u HeroBaseService) UpdateHeroBaseData(c *gin.Context) (err error) {
	type herolist struct {
		HeroId    int    `form:"heroId" json:"heroId" binding:"required"`
		Name      string `form:"name" json:"name" binding:"required"`
		Position1 string `form:"position1" json:"position1" binding:"required"`
		Position2 string `form:"position2" json:"position2" binding:"required"`
	}

	var form struct {
		Heros []herolist `form:"heros" json:"heros" binding:"required"`
	}
	if err := c.ShouldBind(&form); err != nil {
		return common.ErrNew(err, common.ParamErr)
	}

	for i := 0; i < len(form.Heros); i++ {
		herobase := &model.HeroBase{
			Name: form.Heros[i].Name,
		}
		if err := model.DB.Where(&herobase).First(&herobase).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			herobase = &model.HeroBase{
				HeroId:    form.Heros[i].HeroId,
				Name:      form.Heros[i].Name,
				Position1: form.Heros[i].Position1,
				Position2: form.Heros[i].Position2,
			}
			if err := model.DB.Create(&herobase).Error; err != nil {
				// return common.ErrNew(err, common.SysErr)
				return errors.New("创建新英雄基础数据失败")
			}
		} else if err == nil {
			herobase = &model.HeroBase{
				HeroId:    form.Heros[i].HeroId,
				Name:      form.Heros[i].Name,
				Position1: form.Heros[i].Position1,
				Position2: form.Heros[i].Position2,
			}
			if err := model.DB.Where("name = ?", form.Heros[i].Name).Select("hero_id", "name", "position1", "position2").Updates(&herobase).Error; err != nil {
				// return common.ErrNew(err, common.SysErr)
				return errors.New("更新英雄基础数据失败")
			}
		} else {
			// return common.ErrNew(err, common.SysErr)
			return errors.New("查询旧英雄基础数据失败")
		}
	}
	return nil
}
