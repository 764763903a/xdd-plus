package controllers

import (
	"github.com/cdle/xdd/models"
)

type AccountController struct {
	BaseController
}

func (c *AccountController) NextPrepare() {
	c.Logined()
}

func (c *AccountController) List() {
	var page = c.GetQueryInt("page")
	var limit = c.GetQueryInt("limit")
	var cks = models.GetJdCookies()
	if !c.Master {
		tmp := cks
		cks = []models.JdCookie{}
		for _, ck := range tmp {
			if ck.PtPin == c.PtPin {
				cks = append(cks, ck)
				break
			}
		}
	}
	var len = len(cks)
	var total = []int{len}
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 1
	}
	var from = (page - 1) * limit
	var to = page * limit
	if from >= len-1 {
		from = len - 1
	}
	if to >= len {
		to = len
	}
	if from < 0 {
		from = 0
	}
	var data = cks[from:to]
	c.Data["json"] = map[string]interface{}{
		"code":    200,
		"data":    data,
		"message": total,
	}
	c.ServeJSON()
}

func (c *AccountController) CreateOrUpdate() {
	ps := &models.JdCookie{}
	c.Validate(ps)
	if ps.PtPin != "" {
		ps.Pool = ""
		if !c.Master {
			ps.Priority = 0
			ps.PtKey = ""
			ps.PtPin = c.PtPin
		}
		ps.Updates(*ps)
	}
	go func() {
		models.Save <- &models.JdCookie{}
	}()
	c.Response(nil, "操作成功")
}

func (c *AccountController) Admin() {
	c.Ctx.WriteString(models.Admin)
}
