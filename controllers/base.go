package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
	"github.com/TheBeege/mentor-me-api/models"
	"fmt"
)

type BaseController struct {
	beego.Controller

	User *models.User
	IsLogin bool
}

type NestPreparer interface {
	NestPrepare()
}

type NestFinisher interface {
	NestFinish()
}

func (c *BaseController) Prepare() {
	c.SetParams()

	c.IsLogin = c.GetSession("userinfo") != nil
	if c.IsLogin {
		c.User = c.GetLogin()
	}

	c.Data["IsLogin"] = c.IsLogin
	c.Data["Userinfo"] = c.User

	c.Data["HeadStyles"] = []string{}
	c.Data["HeadScripts"] = []string{}

	c.Layout = "base.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["BaseHeader"] = "header.tpl"
	c.LayoutSections["BaseFooter"] = "footer.tpl"

	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (c *BaseController) Finish() {
	if app, ok := c.AppController.(NestFinisher); ok {
		app.NestFinish()
	}
}

func (c *BaseController) GetLogin() *models.User {
	u, err := models.GetUserById(c.GetSession("userinfo").(int))
	if err == nil {
		return nil
	}
	return u
}

func (c *BaseController) DelLogin() {
	c.DelSession("userinfo")
}

func (c *BaseController) SetLogin(user *models.User) {
	c.SetSession("userinfo", user.Id)
}

func (c *BaseController) LoginPath() string {
	return c.URLFor("BaseController.Login")
}

func (c *BaseController) SetParams() {
	c.Data["Params"] = make(map[string]string)
	for k, v := range c.Input() {
		c.Data["Params"].(map[string]string)[k] = v[0]
	}
}

func (c *BaseController) BuildRequestUrl(uri string) string {
	if uri == "" {
		uri = c.Ctx.Input.URI()
	}
	return fmt.Sprintf("%s:%s%s",
		c.Ctx.Input.Site(), strconv.FormatInt(int64(c.Ctx.Input.Port()), 10), uri)
}