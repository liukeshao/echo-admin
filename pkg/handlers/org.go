package handlers

import "github.com/liukeshao/echo-admin/pkg/services"

type OrgHandler struct {
}

func NewOrgHandler(c *services.Container) *OrgHandler {
	return &OrgHandler{}
}

func (h *OrgHandler) List() error {
	return nil
}
