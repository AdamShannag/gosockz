package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/AdamShannag/gosockz/examples/in_memory/handler"
	"github.com/AdamShannag/gosockz/examples/in_memory/otp"
	"github.com/AdamShannag/gosockz/examples/in_memory/repo"
	"github.com/AdamShannag/gosockz/examples/in_memory/util"
	"github.com/AdamShannag/gosockz/manager"
	"github.com/AdamShannag/gosockz/router"
)

var userStore = util.NewInMemoryUserStore(
	util.User{Username: "adam", Password: "adam123"},
	util.User{Username: "yaseen", Password: "yaseen123"},
	util.User{Username: "mike", Password: "mike123"},
)

func main() {

	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()

	upgrader := util.NewWebsocketUpgrader(512, 512, util.CheckOrigin([]string{"http://localhost:8080"}))

	r := router.NewEventRouter().
		Handle(EventBroadcastMessage, BroadcastMessageHandler).
		Handle(EventSendMessage, SendMessageHandler).
		Handle(EventChangeSession, SessionHandler)

	m := manager.NewManager(repo.NewInMemoryClientRepo(), upgrader, r)

	setupAPI(handler.NewHandler(
		m,
		userStore,
		otp.NewRetentionMap(ctx, 5*time.Second),
	))

	log.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func setupAPI(h *handler.Handler) {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/ws", h.ServeWs)
}
