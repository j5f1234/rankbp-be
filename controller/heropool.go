package controller

import (
	"fmt"
	"net/http"
	"rankbp-be/common"

	"github.com/gin-gonic/gin"
)

type HeroPoolController struct {
}

// 修改英雄池
func (u HeroPoolController) ChangeHeropool(c *gin.Context) {
	err := srv.HeroPoolService.ChangeHeropool(c)
	if err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	resp := NullRet{}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

// 发送英雄池数据
func (u HeroPoolController) SendHeropool(c *gin.Context) {
	resp, err := srv.SendHeropool(c)
	if err != nil {
		fmt.Printf("controller %v\n", err)
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp.Heropool))
}
