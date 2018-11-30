package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myapp/models"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Add() {
	if !(checkAccount(c.Ctx)) {
		c.Redirect("/login", 302)
		return
	}
	id := c.Input().Get("tid")
	nickname := c.Input().Get("nickname")
	content := c.Input().Get("content")
	err := models.AddReply(id, nickname, content)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect(fmt.Sprint("/topic/view/", id), 302)
}

func (c *ReplyController) Delete() {
	if !(checkAccount(c.Ctx)) {
		c.Redirect("/login", 302)
		return
	}
	id := c.Ctx.Input.Param("0")
	tid := c.Ctx.Input.Param("1")
	err := models.DeleteReply(id)
	if err != nil {
		return
	}
	c.Redirect("/topic/view/"+tid, 302)
}
