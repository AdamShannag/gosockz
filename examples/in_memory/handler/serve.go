package handler

import (
	"log"
	"net/http"

	"github.com/AdamShannag/gosockz/client"
)

func (h *Handler) ServeWs(w http.ResponseWriter, r *http.Request) {

	otp := r.URL.Query().Get("otp")
	if otp == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !h.otps.VerifyOTP(otp) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	conn, err := h.m.WebsocketUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := client.NewDefaultClient(username, conn, h.m)
	_ = h.m.AddClient(client)

	go client.ReadMessages()
	go client.WriteMessages()

}
