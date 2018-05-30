package controllers

import (
	//	b64 "encoding/base64"
	"github.com/ParallelCoinTeam/duocli/app/routes"
	"github.com/revel/revel"
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

	server := c.Session["server"]
	return c.Render(server)
}
