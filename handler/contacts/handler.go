package contacts

import (
	"github.com/tgo-team/tgo-talkapi/cmd"
	"github.com/tgo-team/tgo-talkapi/config"
	"github.com/tgo-team/tgo-talkapi/ctrl"
)

type Handler struct {
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
	}
}

func (h *Handler) syncContacts(c *cmd.Context)  {

}


func (h *Handler) RegisterHandler(c *ctrl.Controller) {
	// 同步联系人
	c.RegisterHandlerFunc("sync_contacts", h.syncContacts)
}
