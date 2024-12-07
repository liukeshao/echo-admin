package routers

import (
	"github.com/liukeshao/echo-admin/pkg/handlers"
	"github.com/liukeshao/echo-admin/pkg/services"
)

type Endpoints struct {
	container *services.Container

	account *handlers.AccountHandler
	org     *handlers.OrgHandler
}

func NewEndpoints(c *services.Container) *Endpoints {
	e := &Endpoints{
		container: c,
	}
	e.init()
	return e
}

func (e *Endpoints) init() {
	// handlers
	e.account = handlers.NewAccountHandler(e.container)
	e.org = handlers.NewOrgHandler(e.container)
}
