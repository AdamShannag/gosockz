package handler

import (
	"github.com/AdamShannag/gosockz/examples/in_memory/otp"
	"github.com/AdamShannag/gosockz/examples/in_memory/util"
	"github.com/AdamShannag/gosockz/types"
)

type Handler struct {
	m         types.Manager
	userStore *util.InMemoryUserStore
	otps      otp.RetentionMap
}

func NewHandler(m types.Manager, userStore *util.InMemoryUserStore, otps otp.RetentionMap) *Handler {
	return &Handler{
		m, userStore, otps,
	}
}
