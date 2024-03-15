package controller

import (
	"fmt"
	"net/http"
	"rankbp-be/common"

	"github.com/gin-gonic/gin"
)

type HeroBaseController struct {
}

// 更新英雄基础数据
func (u HeroBaseController) UpdateHeroBaseData(c *gin.Context) {
	err := srv.HeroBaseService.UpdateHeroBaseData(c)
	if err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp := NullRet{}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}
