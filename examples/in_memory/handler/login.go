package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	type userLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req userLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !h.userStore.VerifyUser(req.Username, req.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	type response struct {
		OTP string `json:"otp"`
	}

	otp := h.otps.NewOTP()

	resp := response{
		OTP: otp.Key,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
