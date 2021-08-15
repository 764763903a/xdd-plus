package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/cdle/xdd/models"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var validate *validator.Validate
var trans ut.Translator

func init() {
	//验证器注册翻译器
	var zhCh = zh.New()
	validate = validator.New()
	var uni = ut.New(zhCh)
	trans, _ = uni.GetTranslator("zh")
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

//BaseController 基础控制器
type BaseController struct {
	beego.Controller
	PtPin  string
	Master bool
}

//NextPrepare 下一个准备
type NextPrepare interface {
	NextPrepare()
}

//Prepare 准备
func (c *BaseController) Prepare() {
	// c.Ctx.ResponseWriter.Header().Add("Master-IP-Address", models.GetMasteraddr())
	if app, ok := c.AppController.(NextPrepare); ok {
		app.NextPrepare()
	}
}

//Response 响应
func (c *BaseController) Response(ps ...interface{}) { //数据、信息、状态码
	rsp := struct {
		//状态码
		Code int `json:"code"` // 0 成功 1 失败
		//数据
		Data interface{} `json:"data"`
		//描述信息
		Msg string `json:"msg"`
	}{}
	switch len(ps) {
	case 3:
		rsp.Code = ps[2].(int)
		fallthrough
	case 2:
		switch ps[1].(type) {
		case string:
			rsp.Msg = ps[1].(string)
		case error:
			rsp.Msg = ps[1].(error).Error()
		}
		fallthrough
	case 1:
		rsp.Data = ps[0]
	}
	c.Data["json"] = rsp
	c.ServeJSON()
	c.StopRun()
}

//ResponseError 响应错误
func (c *BaseController) ResponseError(ps ...interface{}) *BaseController {
	if ps[0] == nil {
		return c
	}
	// var status = http.StatusBadRequest
	var text = ""

	for _, p := range ps {
		switch t := p.(type) {
		case int: //状态码
			// status = t
			break
		case error: //错误
			text = t.Error()
			break
		case string: //字符描述
			text = t
			break
		}
	}
	// c.Ctx.ResponseWriter.WriteHeader(status)
	// if text != "" {
	// 	c.Ctx.WriteString(text)
	// }
	c.Response(nil, text, 1)
	// c.StopRun()
	return nil
}

//Logined 登录
func (c *BaseController) Logined() *BaseController {
	if models.ExecPath == "/Users/cdle/Desktop/xdd" { //作者调试
		c.Master = true
		return c
	}
	if v := c.GetSession("pin"); v == nil {
		c.Ctx.Redirect(302, "/")
		c.StopRun()
	} else {
		c.PtPin = v.(string)
		if strings.Contains(models.Config.Master, v.(string)) {
			c.Master = true
		}
	}
	return c
}

//Validate 表单验证
func (c *BaseController) Validate(ps interface{}) *BaseController {
	c.ResponseError(json.Unmarshal(c.Ctx.Input.CopyBody(10000000), ps), http.StatusBadRequest)
	if err := validate.Struct(ps); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			c.ResponseError(err.Translate(trans), http.StatusBadRequest)
		}
	}
	return c
}

//GetPathInt64
func (c *BaseController) GetPathInt64(v string) int64 {
	r := c.Ctx.Input.Param(":" + v)
	if r == "" {
		return 0
	}
	i, err := strconv.Atoi(r)
	c.ResponseError(err)
	return int64(i)
}

//GetPathInt
func (c *BaseController) GetPathInt(v string) int {
	r := c.Ctx.Input.Param(":" + v)
	if r == "" {
		return 0
	}
	i, err := strconv.Atoi(r)
	c.ResponseError(err)
	return i
}

//GetPathInt32
func (c *BaseController) GetPathInt32(v string) int32 {
	r := c.Ctx.Input.Param(":" + v)
	if r == "" {
		return 0
	}
	i, err := strconv.Atoi(r)
	c.ResponseError(err)
	return int32(i)
}

//GetQueryInt64
func (c *BaseController) GetQueryInt64(v string) int64 {
	r := c.GetString(v)
	if r == "" {
		return 0
	}
	i, err := strconv.Atoi(r)
	c.ResponseError(err)
	return int64(i)
}

//GetQueryInt
func (c *BaseController) GetQueryInt(v string) int {
	r := c.GetString(v)
	if r == "" {
		return 0
	}
	i, err := strconv.Atoi(r)
	c.ResponseError(err)
	return i
}

//GetQueryInt32
func (c *BaseController) GetQueryInt32(v string) int32 {
	r := c.GetString(v)
	if r == "" {
		return 0
	}
	i, err := strconv.Atoi(r)
	c.ResponseError(err)
	return int32(i)
}
