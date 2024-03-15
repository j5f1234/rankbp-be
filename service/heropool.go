package service

import (
	"errors"
	"rankbp-be/common"
	"rankbp-be/model"

	"github.com/gin-gonic/gin"
)

type HeroPoolService struct {
}

// 修改英雄池
func (u HeroPoolService) ChangeHeropool(c *gin.Context) (err error) {
	type pool struct {
		Best []string `form:"best" json:"best" binding:"required"`
		Good []string `form:"good" json:"good" binding:"required"`
	}
	var poolform struct {
		Heropool []pool `form:"heropool" json:"heropool" binding:"required"`
	}
	if err := c.ShouldBind(&poolform); err != nil {
		return common.ErrNew(err, common.ParamErr)
	}

	position := [5]string{"TOP", "JUNGLE", "MID", "ADC", "SUPPORT"}

	hero := &model.HeroPool{}
	if err := model.DB.Where("1=1").Delete(&hero).Error; err != nil {
		return common.ErrNew(err, common.SysErr)
	}

	if err := model.DB.Exec("ALTER TABLE hero_pools AUTO_INCREMENT = 1").Error; err != nil {
		return common.ErrNew(err, common.SysErr)
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < len(poolform.Heropool[i].Best); j++ {
			hero = &model.HeroPool{
				Position: position[i],
				Name:     poolform.Heropool[i].Best[j],
				Skilled:  1,
			}
			if err := model.DB.Create(&hero).Error; err != nil {
				// return common.ErrNew(err, common.SysErr)
				return errors.New("创建高熟练度英雄失败")
			}
		}

		for j := 0; j < len(poolform.Heropool[i].Good); j++ {
			hero = &model.HeroPool{
				Position: position[i],
				Name:     poolform.Heropool[i].Good[j],
				Skilled:  0,
			}
			act := 1
			for k := 0; k < len(poolform.Heropool[i].Best); k++ {
				if poolform.Heropool[i].Best[k] == poolform.Heropool[i].Good[j] {
					act = 0
					break
				}
			}
			if act == 1 {
				if err := model.DB.Create(&hero).Error; err != nil {
					// return common.ErrNew(err, common.SysErr)
					return errors.New("创建低熟练度英雄失败")
				}
			}
		}
	}
	return nil
}

type pool struct {
	Best []string `form:"best" json:"best" binding:"required"`
	Good []string `form:"good" json:"good" binding:"required"`
}
type poolformstruct struct {
	Heropool [5]pool `form:"heropool" json:"heropool" binding:"required"`
}

// 发送英雄池数据
func (u HeroPoolService) SendHeropool(c *gin.Context) (resp poolformstruct, err error) {
	var heropoolresp poolformstruct
	position := [5]string{"TOP", "JUNGLE", "MID", "ADC", "SUPPORT"}
	for i := 0; i < 5; i++ {
		hero := []model.HeroPool{{
			Position: position[i],
		}}

		if err := model.DB.Where("position = ?", position[i]).Find(&hero).Error; err == nil {
			heropoolresp.Heropool[i].Best = make([]string, 0)
			heropoolresp.Heropool[i].Good = make([]string, 0)
			for j := 0; j < len(hero); j++ {
				if hero[j].Skilled == 1 {
					heropoolresp.Heropool[i].Best = append(heropoolresp.Heropool[i].Best, hero[j].Name)
				} else {
					heropoolresp.Heropool[i].Good = append(heropoolresp.Heropool[i].Good, hero[j].Name)
				}
			}
		} else {
			return heropoolresp, common.ErrNew(err, common.SysErr)
		}
	}
	return heropoolresp, nil
}
