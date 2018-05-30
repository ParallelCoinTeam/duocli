package controllers

import (
//	b64 "encoding/base64"
	"github.com/revel/revel"
	"github.com/ParallelCoinTeam/duocli/app/routes"
	"github.com/ParallelCoinTeam/duocli/app/services"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {

	server := c.Request.URL.Query().Get("server")
	if server != "" {
		for k := range c.Session {
			delete(c.Session, k)
		}
		c.Session["server"] = server
	}
	return c.Redirect(routes.Application.About())
}

func (c Application) About() revel.Result {

	user_id := c.Session["user_id"]
	user_name := c.Session["user_name"]
	server := c.Session["server"]
	return c.Render(user_id, user_name, server)
}

func (c Application) CurlTest() revel.Result {

	user_id := c.Session["user_id"]
	user_name := c.Session["user_name"]
	server := c.Session["server"]
	return c.Render(user_id, user_name, server)
}
