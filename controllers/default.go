package controllers

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
	"beegomyblog/models"
)

type MainController struct {
	beego.Controller
}

// @router / [get]
func (c *MainController) Get() {
    posts := models.GetPosts()
    logs.Info(posts)
	c.Data["posts"] = posts
	c.TplName = "index.tpl"
}

// @router /post/:file_name [get]
func (c *MainController) GetPost() {
	file_name := c.Ctx.Input.Param(":file_name")
    post := models.GetPost(file_name)
    logs.Info(post)
	c.Data["post"] = post
	c.TplName = "post.tpl"
}

// @router /aboutme [get]
func (c *MainController) AbountMe() {
    post := models.GetAboutMe()
	c.Data["post"] = post
    c.TplName = "aboutme.tpl"
}
